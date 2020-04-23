// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package spaces

import (
	"fmt"
	"strings"

	"github.com/juju/collections/set"
	"github.com/juju/errors"

	"github.com/juju/juju/core/network"
)

// unitNetwork represents a group of units and the subnets
// to which they are *all* connected.
type unitNetwork struct {
	unitNames set.Strings
	// subnets are those that all the units are connected to.
	// Not that these are populated from the *current* network topology.
	subnets network.SubnetInfos
}

// hasSameConnectivity returns true if the input set of subnets
// matches the unitNetwork subnets exactly, based on ID.
func (n *unitNetwork) hasSameConnectivity(subnets network.SubnetInfos) bool {
	return n.subnets.EqualTo(subnets)
}

// isConnectedTo returns true if the unitNetwork has a subnet in common with
// the input space.
// Note that we compare subnet IDs and not space IDs, because the subnets are
// populated from the existing topology, whereas the input space comes from
// the hypothetical new topology where one or more subnets
// will have changed space.
func (n *unitNetwork) isConnectedTo(space network.SpaceInfo) bool {
	spaceSubs := space.Subnets
	for _, sub := range n.subnets {
		if spaceSubs.ContainsID(sub.ID) {
			return true
		}
	}
	return false
}

// affectedNetworks groups unique unit networks by application name.
// It facilitates checking whether the connectedness of application units
// is able to honour changing space topology based on application
// constraints and endpoint bindings.
type affectedNetworks struct {
	// subnets identifies the subnets that are being moved.
	subnets network.IDSet
	// newSpace is the name of the space that the subnets are being moved to.
	newSpace string
	// spaces is a the target space topology.
	spaces network.SpaceInfos
	// appNetworks are are all unit subnet connectivity grouped by application
	// for any that may be affected by moving the subnets above.
	appNetworks map[string][]unitNetwork
	force       bool
}

// newAffectedNetworks returns a new affectedNetworks reference for
// verification of the movement of the input subnets to the input space.
// The input space topology is manipulated to represent the topology that
// would result from the move.
func newAffectedNetworks(
	movingSubnets network.IDSet, spaceName string, currentTopology network.SpaceInfos, force bool,
) (*affectedNetworks, error) {

	// We need to indicate that any moving fan underlays include
	// their overlays as being affected by a move.
	movingOverlays, err := currentTopology.FanOverlaysFor(movingSubnets)
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, overlay := range movingOverlays {
		movingSubnets.Add(overlay.ID)
	}

	// Now get the topology as would result from moving all of these subnets.
	newTopology, err := currentTopology.MoveSubnets(movingSubnets, spaceName)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &affectedNetworks{
		subnets:     movingSubnets,
		newSpace:    spaceName,
		spaces:      newTopology,
		appNetworks: make(map[string][]unitNetwork),
		force:       force,
	}, nil
}

// processMachines iterates over the input machines,
// looking at the subnets they are connected to.
// Any machines connected to a moving subnet have their unit networks
// included for for later verification.
func (n *affectedNetworks) processMachines(machines []Machine) error {
	for _, machine := range machines {
		addresses, err := machine.AllAddresses()
		if err != nil {
			return errors.Trace(err)
		}

		var includesMover bool
		var machineSubnets network.SubnetInfos
		for _, address := range addresses {
			// TODO (manadart 2020-04-22): Note that the Subnet method here
			// looks up the subnet info based on the address CIDR.
			// The backing schema for this needs to be rethought for multi-net
			// capability.
			sub, err := address.Subnet()
			if err != nil {
				return errors.Trace(err)
			}
			machineSubnets = append(machineSubnets, sub)

			if n.subnets.Contains(sub.ID) {
				includesMover = true
			}
		}

		// We only consider this machine if it has an
		// address in one of the moving subnets.
		if includesMover {
			if err = n.includeMachine(machine, machineSubnets); err != nil {
				return errors.Trace(err)
			}
		}
	}

	return nil
}

// includeMachine ensures that the units on the machine and their collection
// of subnet connectedness are included as affectedNetworks to be validated.
func (n *affectedNetworks) includeMachine(machine Machine, subnets network.SubnetInfos) error {
	units, err := machine.Units()
	if err != nil {
		return errors.Trace(err)
	}

	for _, unit := range units {
		appName := unit.ApplicationName()
		unitNets, ok := n.appNetworks[appName]
		if !ok {
			n.appNetworks[appName] = []unitNetwork{}
		}

		var present bool

		for _, unitNet := range unitNets {
			if unitNet.hasSameConnectivity(subnets) {
				unitNet.unitNames.Add(unit.Name())
				present = true
			}
		}

		if !present {
			n.appNetworks[appName] = append(unitNets, unitNetwork{
				unitNames: set.NewStrings(unit.Name()),
				subnets:   subnets,
			})
		}
	}

	return nil
}

// ensureConstraintIntegrity checks that moving subnets to the new space does
// not violate any application space constraints.
// If force is true, violations are logged as warnings,
// otherwise an error is returned.
func (n *affectedNetworks) ensureConstraintIntegrity(cons map[string]set.Strings) error {
	for appName := range n.appNetworks {
		// We know there are always constraints; this is ensured by the caller.
		spaces := cons[appName]

		if err := n.ensureNegativeConstraintIntegrity(appName, spaces); err != nil {
			return errors.Trace(err)
		}

		if err := n.ensurePositiveConstraintIntegrity(appName, spaces); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

// ensureNegativeConstraintIntegrity checks that the input application does not
// have a negative space constraint for the proposed destination space.
func (n *affectedNetworks) ensureNegativeConstraintIntegrity(appName string, spaceConstraints set.Strings) error {
	if spaceConstraints.Contains("^" + n.newSpace) {
		msg := fmt.Sprintf("moving subnet(s) to space %q violates space constraints "+
			"for application %q: %s", n.newSpace, appName, strings.Join(spaceConstraints.SortedValues(), ", "))

		if !n.force {
			return errors.New(msg)
		}
		logger.Warningf(msg)
	}

	return nil
}

// ensurePositiveConstraintIntegrity checks that for each positive space
// constraint, comparing the input application's unit subnet connectivity to
// the target topology determines the constraint to be satisfied.
func (n *affectedNetworks) ensurePositiveConstraintIntegrity(appName string, spaceConstraints set.Strings) error {
	unitNets := n.appNetworks[appName]

	for _, spaceName := range spaceConstraints.Values() {
		if strings.HasPrefix(spaceName, "^") {
			continue
		}

		conSpace := n.spaces.GetByName(spaceName)
		if conSpace == nil {
			return errors.NotFoundf("space with name %q", spaceName)
		}

		for _, unitNet := range unitNets {
			if !unitNet.isConnectedTo(*conSpace) {
				msg := fmt.Sprintf(
					"moving subnet(s) to space %q violates space constraints "+
						"for application %q: %s\n\tunits not connected to the space: %s",
					n.newSpace,
					appName,
					strings.Join(spaceConstraints.SortedValues(), ", "),
					strings.Join(unitNet.unitNames.SortedValues(), ", "),
				)

				if !n.force {
					return errors.New(msg)
				}
				logger.Warningf(msg)
			}
		}
	}

	return nil
}
