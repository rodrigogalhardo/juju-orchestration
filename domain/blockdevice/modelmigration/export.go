// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelmigration

import (
	"context"

	"github.com/juju/description/v5"
	"github.com/juju/errors"
	"github.com/juju/loggo"

	"github.com/juju/juju/core/modelmigration"
	"github.com/juju/juju/domain/blockdevice"
	"github.com/juju/juju/domain/blockdevice/service"
	"github.com/juju/juju/domain/blockdevice/state"
)

var logger = loggo.GetLogger("juju.migration.blockdevices")

// RegisterExport registers the export operations with the given coordinator.
func RegisterExport(coordinator Coordinator) {
	coordinator.Add(&exportOperation{})
}

// ExportService provides a subset of the block device domain
// service methods needed for block device export.
type ExportService interface {
	AllBlockDevices(ctx context.Context) (map[string]blockdevice.BlockDevice, error)
}

// exportOperation describes a way to execute a migration for
// exporting block devices.
type exportOperation struct {
	modelmigration.BaseOperation

	service ExportService
}

// Setup implements Operation.
func (e *exportOperation) Setup(scope modelmigration.Scope) error {
	// We must not use a watcher during migration, so it's safe to pass a
	// nil watcher factory.
	e.service = service.NewService(
		state.NewState(scope.ModelDB()), nil, logger)
	return nil
}

// Execute the export, adding the block devices to the model.
func (e *exportOperation) Execute(ctx context.Context, model description.Model) error {
	blockDevices, err := e.service.AllBlockDevices(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	for machineId, bd := range blockDevices {
		err := model.AddBlockDevice(machineId, description.BlockDeviceArgs{
			Name:           bd.DeviceName,
			Links:          bd.DeviceLinks,
			Label:          bd.Label,
			UUID:           bd.UUID,
			HardwareID:     bd.HardwareId,
			SerialID:       bd.SerialId,
			WWN:            bd.WWN,
			BusAddress:     bd.BusAddress,
			Size:           bd.SizeMiB,
			FilesystemType: bd.FilesystemType,
			InUse:          bd.InUse,
			MountPoint:     bd.MountPoint,
		})
		if err != nil {
			return errors.Annotatef(err, "adding block device for machine %q", machineId)
		}
	}
	return nil
}