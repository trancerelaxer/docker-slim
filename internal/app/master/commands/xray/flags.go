package xray

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Xray command flag names
const (
	FlagChanges          = "changes"
	FlagLayer            = "layer"
	FlagAddImageManifest = "add-image-manifest"
	FlagAddImageConfig   = "add-image-config"
	FlagLayerChangesMax  = "layer-changes-max"
	FlagAllChangesMax    = "all-changes-max"
	FlagAddChangesMax    = "add-changes-max"
	FlagModifyChangesMax = "modify-changes-max"
	FlagDeleteChangesMax = "delete-changes-max"
	FlagChangePath       = "change-path"
	FlagChangeData       = "change-data"
)

// Xray command flag usage info
const (
	FlagChangesUsage          = "Show layer change details for the selected change type (values: none, all, delete, modify, add)"
	FlagLayerUsage            = "Show details for the selected layer (using layer index or ID)"
	FlagAddImageManifestUsage = "Add raw image manifest to the command execution report file"
	FlagAddImageConfigUsage   = "Add raw image config object to the command execution report file"
	FlagLayerChangesMaxUsage  = "Maximum number of changes to show for each layer"
	FlagAllChangesMaxUsage    = "Maximum number of changes to show for all layers"
	FlagAddChangesMaxUsage    = "Maximum number of 'add' changes to show for all layers"
	FlagModifyChangesMaxUsage = "Maximum number of 'modify' changes to show for all layers"
	FlagDeleteChangesMaxUsage = "Maximum number of 'delete' changes to show for all layers"
	FlagChangePathUsage       = "Include changes for the files that match the path pattern (Glob/Match in Go and **)"
	FlagChangeDataUsage       = "Include changes for the files that match the data pattern (regex)"
)

var Flags = map[string]cli.Flag{
	FlagChanges: cli.StringSliceFlag{
		Name:   FlagChanges,
		Value:  &cli.StringSlice{""},
		Usage:  FlagChangesUsage,
		EnvVar: "DSLIM_CHANGES",
	},
	FlagLayer: cli.StringSliceFlag{
		Name:   FlagLayer,
		Value:  &cli.StringSlice{},
		Usage:  FlagLayerUsage,
		EnvVar: "DSLIM_LAYER",
	},
	FlagAddImageManifest: cli.BoolFlag{
		Name:   FlagAddImageManifest,
		Usage:  FlagAddImageManifestUsage,
		EnvVar: "DSLIM_XRAY_IMAGE_MANIFEST",
	},
	FlagAddImageConfig: cli.BoolFlag{
		Name:   FlagAddImageConfig,
		Usage:  FlagAddImageConfigUsage,
		EnvVar: "DSLIM_XRAY_IMAGE_CONFIG",
	},
	FlagLayerChangesMax: cli.IntFlag{
		Name:   FlagLayerChangesMax,
		Value:  -1,
		Usage:  FlagLayerChangesMaxUsage,
		EnvVar: "DSLIM_XRAY_LAYER_CHANGES_MAX",
	},
	FlagAllChangesMax: cli.IntFlag{
		Name:   FlagAllChangesMax,
		Value:  -1,
		Usage:  FlagAllChangesMaxUsage,
		EnvVar: "DSLIM_XRAY_ALL_CHANGES_MAX",
	},
	FlagAddChangesMax: cli.IntFlag{
		Name:   FlagAddChangesMax,
		Value:  -1,
		Usage:  FlagAddChangesMaxUsage,
		EnvVar: "DSLIM_XRAY_ADD_CHANGES_MAX",
	},
	FlagModifyChangesMax: cli.IntFlag{
		Name:   FlagModifyChangesMax,
		Value:  -1,
		Usage:  FlagModifyChangesMaxUsage,
		EnvVar: "DSLIM_XRAY_MODIFY_CHANGES_MAX",
	},
	FlagDeleteChangesMax: cli.IntFlag{
		Name:   FlagDeleteChangesMax,
		Value:  -1,
		Usage:  FlagDeleteChangesMaxUsage,
		EnvVar: "DSLIM_XRAY_DELETE_CHANGES_MAX",
	},
	FlagChangePath: cli.StringSliceFlag{
		Name:   FlagChangePath,
		Value:  &cli.StringSlice{},
		Usage:  FlagChangePathUsage,
		EnvVar: "DSLIM_XRAY_CHANGE_PATH",
	},
	FlagChangeData: cli.StringSliceFlag{
		Name:   FlagChangeData,
		Value:  &cli.StringSlice{},
		Usage:  FlagChangeDataUsage,
		EnvVar: "DSLIM_XRAY_CHANGE_DATA",
	},
}

func cflag(name string) cli.Flag {
	cf, ok := Flags[name]
	if !ok {
		log.Fatalf("unknown flag='%s'", name)
	}

	return cf
}
