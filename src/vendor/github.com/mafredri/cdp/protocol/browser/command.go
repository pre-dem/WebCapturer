// Code generated by cdpgen. DO NOT EDIT.

package browser

import (
	"github.com/mafredri/cdp/protocol/target"
)

// GetVersionReply represents the return values for GetVersion in the Browser domain.
type GetVersionReply struct {
	ProtocolVersion string `json:"protocolVersion"` // Protocol version.
	Product         string `json:"product"`         // Product name.
	Revision        string `json:"revision"`        // Product revision.
	UserAgent       string `json:"userAgent"`       // User-Agent.
	JsVersion       string `json:"jsVersion"`       // V8 version.
}

// GetWindowBoundsArgs represents the arguments for GetWindowBounds in the Browser domain.
type GetWindowBoundsArgs struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
}

// NewGetWindowBoundsArgs initializes GetWindowBoundsArgs with the required arguments.
func NewGetWindowBoundsArgs(windowID WindowID) *GetWindowBoundsArgs {
	args := new(GetWindowBoundsArgs)
	args.WindowID = windowID
	return args
}

// GetWindowBoundsReply represents the return values for GetWindowBounds in the Browser domain.
type GetWindowBoundsReply struct {
	Bounds Bounds `json:"bounds"` // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// GetWindowForTargetArgs represents the arguments for GetWindowForTarget in the Browser domain.
type GetWindowForTargetArgs struct {
	TargetID target.ID `json:"targetId"` // Devtools agent host id.
}

// NewGetWindowForTargetArgs initializes GetWindowForTargetArgs with the required arguments.
func NewGetWindowForTargetArgs(targetID target.ID) *GetWindowForTargetArgs {
	args := new(GetWindowForTargetArgs)
	args.TargetID = targetID
	return args
}

// GetWindowForTargetReply represents the return values for GetWindowForTarget in the Browser domain.
type GetWindowForTargetReply struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
	Bounds   Bounds   `json:"bounds"`   // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// SetWindowBoundsArgs represents the arguments for SetWindowBounds in the Browser domain.
type SetWindowBoundsArgs struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
	Bounds   Bounds   `json:"bounds"`   // New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
}

// NewSetWindowBoundsArgs initializes SetWindowBoundsArgs with the required arguments.
func NewSetWindowBoundsArgs(windowID WindowID, bounds Bounds) *SetWindowBoundsArgs {
	args := new(SetWindowBoundsArgs)
	args.WindowID = windowID
	args.Bounds = bounds
	return args
}
