//go:build windows

package edge

type Version struct {
	SDKVersion     string
	ReleaseNotes   string
	RuntimeVersion string
	Notes          string
}

var versionMapping = map[string]Version{
	"1.0.2592.51": {
		SDKVersion:     "1.0.2592.51",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10259251",
		RuntimeVersion: "126.0.2592.51",
		Notes: `

<!-- ------------------------------ -->
#### Promotions

No additional APIs have been promoted to Stable and added in this Release SDK.


<!-- ------------------------------ -->
#### Bug fixes

###### Runtime-only

* Disabled 'BreakoutBoxPreferCaptureTimestampInVideoFrame' for WebView2 'TextureStream'.
* Fixed a regression where the 'WindowCloseRequested' event only fires for first 'window.close()' call.
* Fixed a regression where typed arrays in WinRT JavaScript projection could not be handled as 'IDispatch' in the host. 
* Fixed a bug where the autofill popup dismisses immediately and causes a focus change.
* Fixed a bug where WebView2 fails to load because of 'AppPolicyGetWindowingModel'.  ([Issue #4591](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4591))

<!-- end of Jun 2024 Release SDK -->


<!-- ====================================================================== -->
Release Date: June 19, 2024`,
	},
	"1.0.2646-prerelease": {
		SDKVersion:     "1.0.2646-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#102646-prerelease",
		RuntimeVersion: "128.0.2646.0",
		Notes: `

<!-- ------------------------------ -->
#### General features

* Added support for C#/WinRT .NET 6+.


<!-- ------------------------------ -->
#### Experimental features

* Introduced the feature flag 'msWebView2EnableDownloadContentInWebResourceResponseReceived', an an Experimental feature (rather than as a Stable feature).  When this flag is enabled, this allows responses of navigations that become downloads to be available in 'WebResourceResponseReceived'.


<!-- ------------------------------ -->
#### Experimental APIs

The following Experimental APIs have been added in this Prerelease SDK.


<!-- ---------- -->
* Added a new 'SaveFileSecurityCheckStarting' event.  As a developer, you can register a handler on this event to get the file path, filename extension, and document origin URI information.  Then you can apply your own rules to do actions such as the following:
   * Allow saving the file without presenting a default security-warning UI about the file-type policy.
   * Cancel the saving.
   * Create your own UI to manage runtime file-type policies.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.SaveFileSecurityCheckStarting Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.savefilesecuritycheckstarting?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

* [CoreWebView2SaveFileSecurityCheckStartingEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.CancelSave Property](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs.cancelsave?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.DocumentOriginUri Property](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs.documentoriginuri?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.FileExtension Property](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs.fileextension?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.FilePath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs.filepath?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.GetDeferral Method](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs.getdeferral?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.SuppressDefaultPolicy Property](/dotnet/api/microsoft.web.webview2.core.corewebview2savefilesecuritycheckstartingeventargs.suppressdefaultpolicy?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)


##### [WinRT/C#](#tab/winrtcsharp)


* 'CoreWebView2' Class:
   * [CoreWebView2.SaveFileSecurityCheckStarting Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#savefilesecuritycheckstarting)

* [CoreWebView2SaveFileSecurityCheckStartingEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2savefilesecuritycheckstartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.CancelSave Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2savefilesecuritycheckstartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#cancelsave)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.DocumentOriginUri Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2savefilesecuritycheckstartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#documentoriginuri)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.FileExtension Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2savefilesecuritycheckstartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#fileextension)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.FilePath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2savefilesecuritycheckstartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#filepath)
   * [CoreWebView2SaveFileSecurityCheckStartingEventArgs.SuppressDefaultPolicy Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2savefilesecuritycheckstartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#suppressdefaultpolicy)

  
##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental27](/microsoft-edge/webview2/reference/win32/icorewebview2experimental27?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental27::add_SaveFileSecurityCheckStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimental27?view=webview2-1.0.2646-prerelease&preserve-view=true#add_savefilesecuritycheckstarting)
   * [ICoreWebView2Experimental27::remove_SaveFileSecurityCheckStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimental27?view=webview2-1.0.2646-prerelease&preserve-view=true#remove_savefilesecuritycheckstarting)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventhandler?view=webview2-1.0.2646-prerelease&preserve-view=true)

* [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs](/microsoft-
edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::get_CancelSave](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_cancelsave)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::put_CancelSave](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#put_cancelsave)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::get_DocumentOriginUri](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_documentoriginuri)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::get_FileExtension](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_fileextension)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::get_FilePath](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_filepath)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::get_SuppressDefaultPolicy](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_suppressdefaultpolicy)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::put_SuppressDefaultPolicy](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#put_suppressdefaultpolicy)
   * [ICoreWebView2ExperimentalSaveFileSecurityCheckStartingEventArgs::GetDeferral](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsavefilesecuritycheckstartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#getdeferral)
---


<!-- ---------- -->
* Added a new 'ScreenCaptureStarting' event.  This event is raised whenever the WebView2 and/or iframe that corresponds to the 'CoreWebView2Frame' (or to any of its descendant iframes) requests permission to use the Screen Capture API before the UI is shown.  As a developer, you can then choose to block the UI from being displayed, or allow the UI to be displayed.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.ScreenCaptureStarting Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.screencapturestarting?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

* 'CoreWebView2Frame' Class:
   * [CoreWebView2Frame.ScreenCaptureStarting Event](/dotnet/api/microsoft.web.webview2.core.corewebview2frame.screencapturestarting?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

* 'CoreWebView2ScreenCaptureStartingEventArgs' Class:
   * [CoreWebView2ScreenCaptureStartingEventArgs.Cancel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2screencapturestartingeventargs.cancel?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2ScreenCaptureStartingEventArgs.GetDeferral Method](/dotnet/api/microsoft.web.webview2.core.corewebview2screencapturestartingeventargs.getdeferral?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2ScreenCaptureStartingEventArgs.Handled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2screencapturestartingeventargs.handled?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2ScreenCaptureStartingEventArgs.OriginalSourceFrameInfo Property](/dotnet/api/microsoft.web.webview2.core.corewebview2screencapturestartingeventargs.originalsourceframeinfo?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.ScreenCaptureStarting Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#screencapturestarting)

* 'CoreWebView2Frame' Class:
   * [CoreWebView2Frame.ScreenCaptureStarting Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frame?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#screencapturestarting)

* 'CoreWebView2ScreenCaptureStartingEventArgs' Class:
   * [CoreWebView2ScreenCaptureStartingEventArgs.Cancel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2screencapturestartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#cancel)
   * [CoreWebView2ScreenCaptureStartingEventArgs.GetDeferral Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2screencapturestartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#getdeferral)
   * [CoreWebView2ScreenCaptureStartingEventArgs.Handled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2screencapturestartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#handled)
   * [CoreWebView2ScreenCaptureStartingEventArgs.OriginalSourceFrameInfo Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2screencapturestartingeventargs?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#originalsourceframeinfo)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental26](/microsoft-edge/webview2/reference/win32/icorewebview2experimental26?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental26::add_ScreenCaptureStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimental26?view=webview2-1.0.2646-prerelease&preserve-view=true#add_screencapturestarting)
   * [ICoreWebView2Experimental26::remove_ScreenCaptureStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimental26?view=webview2-1.0.2646-prerelease&preserve-view=true#remove_screencapturestarting)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventhandler?view=webview2-1.0.2646-prerelease&preserve-view=true)

* [ICoreWebView2ExperimentalFrame6](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe6?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalFrame6::add_ScreenCaptureStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe6?view=webview2-1.0.2646-prerelease&preserve-view=true#add_screencapturestarting)
   * [ICoreWebView2ExperimentalFrame6::remove_ScreenCaptureStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe6?view=webview2-1.0.2646-prerelease&preserve-view=true#remove_screencapturestarting)
   * [ICoreWebView2ExperimentalFrameScreenCaptureStartingEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframescreencapturestartingeventhandler?view=webview2-1.0.2646-prerelease&preserve-view=true)

* [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::get_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_cancel)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::put_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#put_cancel)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::get_Handled](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_handled)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::put_Handled](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#put_handled)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::get_OriginalSourceFrameInfo](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#get_originalsourceframeinfo)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::GetDeferral](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#getdeferral)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::Invoke](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#invoke)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::Invoke](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#invoke)
   * [ICoreWebView2ExperimentalScreenCaptureStartingEventArgs::Invoke](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscreencapturestartingeventargs?view=webview2-1.0.2646-prerelease&preserve-view=true#invoke)

---


<!-- ---------- -->
* Added a new 'GetComICoreWebView2' method to the 'CoreWebView2' .NET class that enables you to convert a 'CoreWebView2' between .NET and COM.  Added a new WinRT interface that enables you to convert a 'CoreWebView2' between WinRT and COM.  This allows you to interoperate between libraries that are written in different languages.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.GetComICoreWebView2 Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.getcomicorewebview2?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'ICoreWebView2Interop2' Interface:
   * [ICoreWebView2Interop2::GetComICoreWebView2 Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2interop2?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#getcomicorewebview2)

##### [Win32/C++](#tab/win32cpp)

Not applicable.

---


<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ---------- -->
* Updated the WebMessageObjects API to allow injecting DOM objects into WebView2 content that's constructed via the app, and via the 'CoreWebView2.PostWebMessage' API in the other direction.  Added a new web object type to represent a file system handle that can be posted to the web content to provide it with filesystem access.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.PostWebMessageAsJsonWithAdditionalObjects Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.postwebmessageasjsonwithadditionalobjects?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

* 'CoreWebView2Environment' Class:
   * [CoreWebView2Environment.CreateWebFileSystemDirectoryHandle Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createwebfilesystemdirectoryhandle?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2Environment.CreateWebFileSystemFileHandle Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createwebfilesystemfilehandle?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

* 'CoreWebView2FileSystemHandle' Class:
   * [CoreWebView2FileSystemHandle.Kind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandle.kind?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2FileSystemHandle.Path Property](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandle.path?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * [CoreWebView2FileSystemHandle.Permission Property](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandle.permission?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)

* [CoreWebView2FileSystemHandleKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandlekind?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * 'File'
   * 'Directory'

* [CoreWebView2FileSystemHandlePermission Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandlepermission?view=webview2-dotnet-1.0.2646-prerelease&preserve-view=true)
   * 'ReadOnly'
   * 'ReadWrite'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.PostWebMessageAsJsonWithAdditionalObjects Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#postwebmessageasjsonwithadditionalobjects)

* 'CoreWebView2Environment' Class:
   * [CoreWebView2Environment.CreateWebFileSystemDirectoryHandle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#createwebfilesystemdirectoryhandle)
   * [CoreWebView2Environment.CreateWebFileSystemFileHandle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#createwebfilesystemfilehandle)

* 'CoreWebView2FileSystemHandle' Class:
   * [CoreWebView2FileSystemHandle.Kind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandle?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#kind)
   * [CoreWebView2FileSystemHandle.Path Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandle?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#path)
   * [CoreWebView2FileSystemHandle.Permission Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandle?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true#permission)

* [CoreWebView2FileSystemHandleKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandlekind?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true)
   * 'File'
   * 'Directory'

* [CoreWebView2FileSystemHandlePermission Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandlepermission?view=webview2-winrt-1.0.2646-prerelease&preserve-view=true)
   * 'ReadOnly'
   * 'ReadWrite'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_23](/microsoft-edge/webview2/reference/win32/icorewebview2_23?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2_23::PostWebMessageAsJsonWithAdditionalObjects](/microsoft-edge/webview2/reference/win32/icorewebview2_23?view=webview2-1.0.2646-prerelease&preserve-view=true#postwebmessageasjsonwithadditionalobjects)

* [ICoreWebView2Environment14](/microsoft-edge/webview2/reference/win32/icorewebview2environment14?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2Environment14::CreateObjectCollection](/microsoft-edge/webview2/reference/win32/icorewebview2environment14?view=webview2-1.0.2646-prerelease&preserve-view=true#createobjectcollection)<!--win32 only-->
   * [ICoreWebView2Environment14::CreateWebFileSystemDirectoryHandle](/microsoft-edge/webview2/reference/win32/icorewebview2environment14?view=webview2-1.0.2646-prerelease&preserve-view=true#createwebfilesystemdirectoryhandle)
   * [ICoreWebView2Environment14::CreateWebFileSystemFileHandle](/microsoft-edge/webview2/reference/win32/icorewebview2environment14?view=webview2-1.0.2646-prerelease&preserve-view=true#createwebfilesystemfilehandle)

* [ICoreWebView2FileSystemHandle](/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2FileSystemHandle::get_Kind](/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle?view=webview2-1.0.2646-prerelease&preserve-view=true#get_kind)
   * [ICoreWebView2FileSystemHandle::get_Path](/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle?view=webview2-1.0.2646-prerelease&preserve-view=true#get_path)
   * [ICoreWebView2FileSystemHandle::get_Permission](/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle?view=webview2-1.0.2646-prerelease&preserve-view=true#get_permission)

* [ICoreWebView2ObjectCollection](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollection?view=webview2-1.0.2646-prerelease&preserve-view=true)
   * [ICoreWebView2ObjectCollection::RemoveValueAtIndex](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollection?view=webview2-1.0.2646-prerelease&preserve-view=true#removevalueatindex)
   * [ICoreWebView2ObjectCollection::InsertValueAtIndex](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollection?view=webview2-1.0.2646-prerelease&preserve-view=true#insertvalueatindex)
   * [ICoreWebView2ObjectCollection::get_Kind](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollection?view=webview2-1.0.2646-prerelease&preserve-view=true#get_kind)

* [COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2646-prerelease&preserve-view=true#corewebview2_file_system_handle_kind)
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND_FILE'
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND_DIRECTORY'

* [COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2646-prerelease&preserve-view=true#corewebview2_file_system_handle_permission)
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_ONLY'
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_WRITE'

---


<!-- ------------------------------ -->
#### Bug fixes

###### Runtime-only

* Fixed a bug in owned-window activation logic for visual hosting.

<!-- end of Jun 2024 Prerelease SDK -->


<!-- ====================================================================== -->
Release Date: May 28, 2024`,
	},
	"1.0.2535.41": {
		SDKVersion:     "1.0.2535.41",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10253541",
		RuntimeVersion: "125.0.2535.41",
		Notes: `

<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ---------- -->
* Support for the Fluent Style Overlay Scrollbar.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ScrollBarStyle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.scrollbarstyle?view=webview2-dotnet-1.0.2535.41&preserve-view=true)

* [CoreWebView2ScrollbarStyle Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2scrollbarstyle?view=webview2-dotnet-1.0.2535.41&preserve-view=true)
   * 'Default'
   * 'FluentOverlay'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ScrollBarStyle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2535.41&preserve-view=true#scrollbarstyle)

* [CoreWebView2ScrollbarStyle Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2scrollbarstyle?view=webview2-winrt-1.0.2535.41&preserve-view=true)
   * 'Default'
   * 'FluentOverlay'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions8](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8?view=webview2-1.0.2535.41&preserve-view=true)
  * [ICoreWebView2EnvironmentOptions8::get_ScrollBarStyle](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8?view=webview2-1.0.2535.41&preserve-view=true#get_scrollbarstyle)
  * [ICoreWebView2EnvironmentOptions8::put_ScrollBarStyle](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8?view=webview2-1.0.2535.41&preserve-view=true#put_scrollbarstyle)

* [COREWEBVIEW2_SCROLLBAR_STYLE enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2535.41&preserve-view=true#corewebview2_scrollbar_style)
   * 'COREWEBVIEW2_SCROLLBAR_STYLE_DEFAULT'
   * 'COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY'

---


<!-- ------------------------------ -->
#### Bug fixes


<!-- ---------- -->
###### Runtime-only

* Fixed a bug where if the 'LaunchingExternalURIScheme' event handler is attached, and the **always remember** checkbox is enabled, and the user selects this checkbox, the dialog is incorrectly displayed again. 
* Fixed an issue where text edit controls in visual hosting would duplicate IME input when losing and then regaining focus.
* Fixed an issue where full-trust UWP apps couldn't display owned windows.


<!-- ---------- -->
###### SDK-only

* Fixed an issue in the SDK causing erroneous \<Platform\> values in the .NET project platforms list.  ([Issue #1755](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1755))

<!-- end of May 2024 Release SDK -->


<!-- ====================================================================== -->
Release Date: May 28, 2024`,
	},
	"1.0.2584-prerelease": {
		SDKVersion:     "1.0.2584-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#102584-prerelease",
		RuntimeVersion: "126.0.2584.0",
		Notes: `

<!-- ------------------------------ -->
#### Experimental features

* Introduced an option to cancel the initial navigation in WebView2, to improve startup performance.  This change is disabled by default, and can be enabled by using the 'msWebView2CancelInitialNavigation' feature flag.


<!-- ------------------------------ -->
#### Experimental APIs

No Experimental APIs have been added in this Prerelease SDK.


<!-- ------------------------------ -->
#### Promotions

No APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
#### Bug fixes


<!-- ---------- -->
###### Runtime and SDK

* Fixed a crash when .NET host object async methods return a null result.  ([Issue #4509](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4509))


<!-- ---------- -->
###### Runtime-only

* Fixed a WebView2 memory leak issue when the window is closed.  ([Issue #4286](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4286))
* Fixed an issue where 'ignoreMemberNotFoundError' wasn't working for .NET objects.  ([Issue #4497](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4497))
* Now returns a proper error code when 'CreateSharedBuffer' is called with 0 buffer size.  ([Issue #4554](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4554))
* Fixed an activation issue for the caret browsing dialog.
* Fixed an issue where the WebView2 Visual Hosting 'CursorChanged' event wasn't firing for custom cursors.

<!-- end of May 2024 Prerelease SDK -->


<!-- ====================================================================== -->
Release Date: April 22, 2024`,
	},
	"1.0.2478.35": {
		SDKVersion:     "1.0.2478.35",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10247835",
		RuntimeVersion: "124.0.2478.35",
		Notes: `

<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ---------- -->
* Added the Runtime selection feature to support more prerelease testing and flighting scenarios.  You can specify 'ReleaseChannels' to choose which channels are searched for during environment creation, and 'ChannelSearchKind' to select a search order.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ChannelSearchKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.channelsearchkind?view=webview2-dotnet-1.0.2478.35&preserve-view=true)
   * [CoreWebView2EnvironmentOptions.ReleaseChannels Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.releasechannels?view=webview2-dotnet-1.0.2478.35&preserve-view=true)

* [CoreWebView2ChannelSearchKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2channelsearchkind?view=webview2-dotnet-1.0.2478.35&preserve-view=true)
   * 'MostStable'
   * 'LeastStable'

* [CoreWebView2ReleaseChannels Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2releasechannels?view=webview2-dotnet-1.0.2478.35&preserve-view=true)
   * 'None'
   * 'Stable'
   * 'Beta'
   * 'Dev'
   * 'Canary'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ChannelSearchKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2478.35&preserve-view=true#channelsearchkind)
   * [CoreWebView2EnvironmentOptions.ReleaseChannels Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2478.35&preserve-view=true#releasechannels)

* [CoreWebView2ChannelSearchKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2channelsearchkind?view=webview2-winrt-1.0.2478.35&preserve-view=true)
   * 'MostStable'
   * 'LeastStable'

* [CoreWebView2ReleaseChannels Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2releasechannels?view=webview2-winrt-1.0.2478.35&preserve-view=true)
   * 'None'
   * 'Stable'
   * 'Beta'
   * 'Dev'
   * 'Canary'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions7](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2478.35&preserve-view=true)
   * [ICoreWebView2EnvironmentOptions7::get_ChannelSearchKind](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2478.35&preserve-view=true#get_channelsearchkind)
   * [ICoreWebView2EnvironmentOptions7::put_ChannelSearchKind](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2478.35&preserve-view=true#put_channelsearchkind)
   * [ICoreWebView2EnvironmentOptions7::get_ReleaseChannels](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2478.35&preserve-view=true#get_releasechannels)
   * [ICoreWebView2EnvironmentOptions7::put_ReleaseChannels](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2478.35&preserve-view=true#put_releasechannels)

* [COREWEBVIEW2_CHANNEL_SEARCH_KIND enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2478.35&preserve-view=true#corewebview2_channel_search_kind)
   * 'COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE'
   * 'COREWEBVIEW2_CHANNEL_SEARCH_KIND_LEAST_STABLE'

* [COREWEBVIEW2_RELEASE_CHANNELS enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2478.35&preserve-view=true#corewebview2_release_channels)
   * 'COREWEBVIEW2_RELEASE_CHANNELS_NONE'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_STABLE'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_BETA'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_DEV'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_CANARY'

---


<!-- ------------------------------ -->
#### Bug fixes


<!-- ---------- -->
###### Runtime-only

* Fixes a potential integer overflow that could lead to a crash when using 'AdditionalObjects' in the WebMessage API.

<!-- end of Apr 2024 Release SDK -->


<!-- ====================================================================== -->
Release Date: April 22, 2024`,
	},
	"1.0.2526-prerelease": {
		SDKVersion:     "1.0.2526-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#102526-prerelease",
		RuntimeVersion: "125.0.2526.0",
		Notes: `

<!-- ------------------------------ -->
#### Breaking changes

The minimum .NET Framework version requirement for .NET WebView2, including WPF and WinForms controls, has been updated from .NET Framework 4.5 to .NET Framework 4.6.2.


<!-- ------------------------------ -->
#### Experimental APIs

The following Experimental APIs have been added in this Prerelease SDK.

* Added 'SaveAs' APIs that allow you to programmatically perform the **Save as** operation.  You can use these APIs to block the default **Save as** dialog, and then either save silently, or build your own UI for **Save as**.  These APIs pertain only to the **Save as** dialog, not the **Download** dialog, which continues to use the existing Download APIs.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
  * [CoreWebView2.SaveAsUIShowing Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.saveasuishowing?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2.ShowSaveAsUIAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.showsaveasuiasync?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)

* [CoreWebView2SaveAsKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2saveaskind?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
   * 'Complete'
   * 'Default'
   * 'HtmlOnly'
   * 'SingleFile'

* [CoreWebView2SaveAsUIResult Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuiresult?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
   * 'Cancelled'
   * 'FileAlreadyExists'
   * 'InvalidPath'
   * 'KindNotSupported'
   * 'Success'

* 'CoreWebView2SaveAsUIShowingEventArgs' Class:
  * [CoreWebView2SaveAsUIShowingEventArgs.AllowReplace Property](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.allowreplace?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2SaveAsUIShowingEventArgs.Cancel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.cancel?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2SaveAsUIShowingEventArgs.ContentMimeType Property](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.contentmimetype?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2SaveAsUIShowingEventArgs.GetDeferral Method](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.getdeferral?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2SaveAsUIShowingEventArgs.Kind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.kind?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2SaveAsUIShowingEventArgs.SaveAsFilePath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.saveasfilepath?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
  * [CoreWebView2SaveAsUIShowingEventArgs.SuppressDefaultDialog Property](/dotnet/api/microsoft.web.webview2.core.corewebview2saveasuishowingeventargs.suppressdefaultdialog?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
  * [CoreWebView2.SaveAsUIShowing Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#saveasuishowing)
  * [CoreWebView2.ShowSaveAsUIAsync Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#showsaveasuiasync)

* [CoreWebView2SaveAsKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveaskind?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true)
   * 'Complete'
   * 'Default'
   * 'HtmlOnly'
   * 'SingleFile'

* [CoreWebView2SaveAsUIResult Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuiresult?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true)
   * 'Cancelled'
   * 'FileAlreadyExists'
   * 'InvalidPath'
   * 'KindNotSupported'
   * 'Success'

* 'CoreWebView2SaveAsUIShowingEventArgs' Class:
   * [CoreWebView2SaveAsUIShowingEventArgs.AllowReplace Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#allowreplace)
   * [CoreWebView2SaveAsUIShowingEventArgs.Cancel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#cancel)
   * [CoreWebView2SaveAsUIShowingEventArgs.ContentMimeType Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#contentmimetype)
   * [CoreWebView2SaveAsUIShowingEventArgs.GetDeferral Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#getdeferral)
   * [CoreWebView2SaveAsUIShowingEventArgs.Kind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#kind)
   * [CoreWebView2SaveAsUIShowingEventArgs.SaveAsFilePath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#saveasfilepath)
   * [CoreWebView2SaveAsUIShowingEventArgs.SuppressDefaultDialog Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2saveasuishowingeventargs?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#suppressdefaultdialog)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental25](/microsoft-edge/webview2/reference/win32/icorewebview2experimental25?view=webview2-1.0.2526-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental25::add_SaveAsUIShowing](/microsoft-edge/webview2/reference/win32/icorewebview2experimental25?view=webview2-1.0.2526-prerelease&preserve-view=true#add_saveasuishowing)
   * [ICoreWebView2Experimental25::remove_SaveAsUIShowing](/microsoft-edge/webview2/reference/win32/icorewebview2experimental25?view=webview2-1.0.2526-prerelease&preserve-view=true#remove_saveasuishowing)
   * [ICoreWebView2Experimental25::ShowSaveAsUI](/microsoft-edge/webview2/reference/win32/icorewebview2experimental25?view=webview2-1.0.2526-prerelease&preserve-view=true#showsaveasui)

* [ICoreWebView2ExperimentalShowSaveAsUICompletedHandler](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalshowsaveasuicompletedhandler?view=webview2-1.0.2526-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalShowSaveAsUICompletedHandler::Invoke](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalshowsaveasuicompletedhandler?view=webview2-1.0.2526-prerelease&preserve-view=true#invoke)

* [ICoreWebView2ExperimentalSaveAsUIShowingEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventhandler?view=webview2-1.0.2526-prerelease&preserve-view=true)

* [COREWEBVIEW2_SAVE_AS_KIND enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.2526-prerelease&preserve-view=true#corewebview2_save_as_kind)
   * 'COREWEBVIEW2_SAVE_AS_KIND_DEFAULT'
   * 'COREWEBVIEW2_SAVE_AS_KIND_HTML_ONLY'
   * 'COREWEBVIEW2_SAVE_AS_KIND_SINGLE_FILE'
   * 'COREWEBVIEW2_SAVE_AS_KIND_COMPLETE'

* [COREWEBVIEW2_SAVE_AS_UI_RESULT enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.2526-prerelease&preserve-view=true#corewebview2_save_as_ui_result)
   * 'COREWEBVIEW2_SAVE_AS_UI_RESULT_SUCCESS'
   * 'COREWEBVIEW2_SAVE_AS_UI_RESULT_INVALID_PATH'
   * 'COREWEBVIEW2_SAVE_AS_UI_RESULT_FILE_ALREADY_EXISTS'
   * 'COREWEBVIEW2_SAVE_AS_UI_RESULT_KIND_NOT_SUPPORTED'
   * 'COREWEBVIEW2_SAVE_AS_UI_RESULT_CANCELLED'

* [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::get_AllowReplace](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#get_allowreplace)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::get_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#get_cancel)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::get_ContentMimeType](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#get_contentmimetype)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::get_Kind](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#get_kind)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::get_SaveAsFilePath](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#get_saveasfilepath)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::get_SuppressDefaultDialog](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#get_suppressdefaultdialog)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::GetDeferral](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#getdeferral)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::put_AllowReplace](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#put_allowreplace)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::put_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#put_cancel)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::put_Kind](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#put_kind)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::put_SaveAsFilePath](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#put_saveasfilepath)
   * [ICoreWebView2ExperimentalSaveAsUIShowingEventArgs::put_SuppressDefaultDialog](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsaveasuishowingeventargs?view=webview2-1.0.2526-prerelease&preserve-view=true#put_suppressdefaultdialog)

---


<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ---------- -->
* Support for the Fluent Style Overlay Scrollbar.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ScrollBarStyle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.scrollbarstyle?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)

* [CoreWebView2ScrollbarStyle Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2scrollbarstyle?view=webview2-dotnet-1.0.2526-prerelease&preserve-view=true)
   * 'Default'
   * 'FluentOverlay'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ScrollBarStyle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true#scrollbarstyle)

* [CoreWebView2ScrollbarStyle Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2scrollbarstyle?view=webview2-winrt-1.0.2526-prerelease&preserve-view=true)
   * 'Default'
   * 'FluentOverlay'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions8](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8?view=webview2-1.0.2526-prerelease&preserve-view=true)
  * [ICoreWebView2EnvironmentOptions8::get_ScrollBarStyle](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8?view=webview2-1.0.2526-prerelease&preserve-view=true#get_scrollbarstyle)
  * [ICoreWebView2EnvironmentOptions8::put_ScrollBarStyle](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8?view=webview2-1.0.2526-prerelease&preserve-view=true#put_scrollbarstyle)

* [COREWEBVIEW2_SCROLLBAR_STYLE enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2526-prerelease&preserve-view=true#corewebview2_scrollbar_style)
   * 'COREWEBVIEW2_SCROLLBAR_STYLE_DEFAULT'
   * 'COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY'

---


<!-- ------------------------------ -->
#### Bug fixes


<!-- ---------- -->
###### Runtime and SDK

* Fixed a bug in WinRT JavaScript projection where passing in a typed array resulted in an "Interface Not Supported" error.  ([Issue #3486](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3486))

* Added support for handling 'out' array parameters in WinRT JavaScript projection.


<!-- ---------- -->
###### Runtime-only

* Fixed a bug where the Image Auto-captioning feature was enabled by default.

* Fixed a bug where if the 'LaunchingExternalURIScheme' event handler is attached, if the **always remember** checkbox is enabled and the user selects this checkbox, the dialog will incorrectly be shown again.

* Fixed 'GetNonClientRegionAtPoint' incorrectly returning 'Nowhere' for some points.

* Fixed a bug where the Text Services Framework would disconnect upon dropping a file onto a WebView2 region.

* Fixed a bug where the View Source **Ctrl+U** keyboard shortcut remained enabled when the 'AreDevToolsEnabled' setting was 'false'.

* Fixed a bug where a composable IME was duplicated upon regaining focus.  ([Issue #1610](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1610))

* Ensured that 'devicePixelRatio' is synchronized with custom rasterization scales.  ([Issue #3060](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3060))

* Fixed a race condition when using 'CallDevToolsProtocolMethod' events in 'NewWindowRequested'.  ([Issue #4181](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4181))

* Fixed a crash that can occur in WPF 'TabIntoCore' when the 'Controller' has been destroyed but the user tries to tab into the control (pressing the **Tab** key).  ([Issue #4452](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4452))

* Ensured that spellcheck takes input language with case-insensitive format.

* Made the Language API more robust regarding user input.

* Fixed a bug where the **Save password?** prompt is not displayed.


<!-- ---------- -->
###### SDK-only

* Fixed missing 'AreBrowserExtensionsEnabled' API in WinRT projection.

<!-- end of Apr. 2024 Prerelease SDK -->


<!-- ====================================================================== -->
Release Date: March 25, 2024`,
	},
	"1.0.2420.47": {
		SDKVersion:     "1.0.2420.47",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10242047",
		RuntimeVersion: "123.0.2420.47",
		Notes: `

<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ---------- -->
* Added a new API to provide hit-testing results on the regions that a WebView2 contains.  This API is useful for visually hosted applications that want to handle mouse events on the non-client area of the WebView2 window.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2CompositionController' Class:
   * [CoreWebView2CompositionController.GetNonClientRegionAtPoint Method](/dotnet/api/microsoft.web.webview2.core.corewebview2compositioncontroller.getnonclientregionatpoint?view=webview2-dotnet-1.0.2420.47&preserve-view=true)
   * [CoreWebView2CompositionController.NonClientRegionChanged Event](/dotnet/api/microsoft.web.webview2.core.corewebview2compositioncontroller.nonclientregionchanged?view=webview2-dotnet-1.0.2420.47&preserve-view=true)
   * [CoreWebView2CompositionController.QueryNonClientRegion Method](/dotnet/api/microsoft.web.webview2.core.corewebview2compositioncontroller.querynonclientregion?view=webview2-dotnet-1.0.2420.47&preserve-view=true)

* 'CoreWebView2NonClientRegionChangedEventArgs' Class:
   * [CoreWebView2NonClientRegionChangedEventArgs.RegionKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2nonclientregionchangedeventargs.regionkind?view=webview2-dotnet-1.0.2420.47&preserve-view=true)

* [CoreWebView2NonClientRegionKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2nonclientregionkind?view=webview2-dotnet-1.0.2420.47&preserve-view=true)
   * 'Caption'
   * 'Client'
   * 'Nowhere'

* 'CoreWebView2Settings' Class:
   * [CoreWebView2Settings.IsNonClientRegionSupportEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2settings.isnonclientregionsupportenabled?view=webview2-dotnet-1.0.2420.47&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2CompositionController' Class:
   * [CoreWebView2CompositionController.GetNonClientRegionAtPoint Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2compositioncontroller?view=webview2-winrt-1.0.2420.47&preserve-view=true#getnonclientregionatpoint)
   * [CoreWebView2CompositionController.NonClientRegionChanged Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2compositioncontroller?view=webview2-winrt-1.0.2420.47&preserve-view=true#nonclientregionchanged)
   * [CoreWebView2CompositionController.QueryNonClientRegion Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2compositioncontroller?view=webview2-winrt-1.0.2420.47&preserve-view=true#querynonclientregion)

* 'CoreWebView2NonClientRegionChangedEventArgs' Class:
   * [CoreWebView2NonClientRegionChangedEventArgs.RegionKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2nonclientregionchangedeventargs?view=webview2-winrt-1.0.2420.47&preserve-view=true#regionkind)

* [CoreWebView2NonClientRegionKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2nonclientregionkind?view=webview2-winrt-1.0.2420.47&preserve-view=true)
   * 'Caption'
   * 'Client'
   * 'Nowhere'

* 'CoreWebView2Settings' Class:
   * [CoreWebView2Settings.IsNonClientRegionSupportEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2settings?view=webview2-winrt-1.0.2420.47&preserve-view=true#isnonclientregionsupportenabled)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2CompositionController4](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4?view=webview2-1.0.2420.47&preserve-view=true)
   * [ICoreWebView2CompositionController4::add_NonClientRegionChanged](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4?view=webview2-1.0.2420.47&preserve-view=true#add_nonclientregionchanged)
   * [ICoreWebView2CompositionController4::GetNonClientRegionAtPoint](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4?view=webview2-1.0.2420.47&preserve-view=true#getnonclientregionatpoint)
   * [ICoreWebView2CompositionController4::QueryNonClientRegion](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4?view=webview2-1.0.2420.47&preserve-view=true#querynonclientregion)
   * [ICoreWebView2CompositionController4::remove_NonClientRegionChanged](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4?view=webview2-1.0.2420.47&preserve-view=true#remove_nonclientregionchanged)

* [ICoreWebView2NonClientRegionChangedEventArgs](/microsoft-edge/webview2/reference/win32/icorewebview2nonclientregionchangedeventargs?view=webview2-1.0.2420.47&preserve-view=true)
   * [ICoreWebView2NonClientRegionChangedEventArgs::get_RegionKind](/microsoft-edge/webview2/reference/win32/icorewebview2nonclientregionchangedeventargs?view=webview2-1.0.2420.47&preserve-view=true#get_regionkind)<!--no put-->

* [ICoreWebView2NonClientRegionChangedEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2nonclientregionchangedeventhandler?view=webview2-1.0.2420.47&preserve-view=true)<!-- Win32-only -->

* [ICoreWebView2RegionRectCollectionView](/microsoft-edge/webview2/reference/win32/icorewebview2regionrectcollectionview?view=webview2-1.0.2420.47&preserve-view=true)<!-- Win32-only -->

* [ICoreWebView2Settings9](/microsoft-edge/webview2/reference/win32/icorewebview2settings9?view=webview2-1.0.2420.47&preserve-view=true)
   * [ICoreWebView2Settings9::get_IsNonClientRegionSupportEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2settings9?view=webview2-1.0.2420.47&preserve-view=true#get_isnonclientregionsupportenabled)
   * [ICoreWebView2Settings9::put_IsNonClientRegionSupportEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2settings9?view=webview2-1.0.2420.47&preserve-view=true#put_isnonclientregionsupportenabled)

* [COREWEBVIEW2_NON_CLIENT_REGION_KIND enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2420.47&preserve-view=true#corewebview2_non_client_region_kind)
   * 'COREWEBVIEW2_NON_CLIENT_REGION_KIND_CAPTION'
   * 'COREWEBVIEW2_NON_CLIENT_REGION_KIND_CLIENT'
   * 'COREWEBVIEW2_NON_CLIENT_REGION_KIND_NOWHERE'

---


<!-- ---------- -->
* Added the 'FailureSourceModulePath' property to the 'ProcessFailedEventArgs' type, to specify the full path of the module that caused the crash in cases of Windows code integrity failures - that is, when a process exited with 'STATUS_INVALID_IMAGE_HASH'.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2ProcessFailedEventArgs' Class:
    * [CoreWebView2ProcessFailedEventArgs.FailureSourceModulePath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2processfailedeventargs.failuresourcemodulepath?view=webview2-dotnet-1.0.2420.47&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2ProcessFailedEventArgs' Class:
    * [CoreWebView2ProcessFailedEventArgs.FailureSourceModulePath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2processfailedeventargs?view=webview2-winrt-1.0.2420.47&preserve-view=true#failuresourcemodulepath)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ProcessFailedEventArgs3](/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs3?view=webview2-1.0.2420.47&preserve-view=true)
    * [ICoreWebView2ProcessFailedEventArgs3::get_FailureSourceModulePath](/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs3?view=webview2-1.0.2420.47&preserve-view=true#get_failuresourcemodulepath)

---


<!-- ------------------------------ -->
#### Bug fixes


###### SDK-only

* The .NET assemblies for WinForms and WPF are now shipped with optimization enabled.  ([Issue #4409](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4409))

<!-- end of Mar. 2024 Release SDK -->


<!-- ====================================================================== -->
Release Date: March 25, 2024`,
	},
	"1.0.2470-prerelease": {
		SDKVersion:     "1.0.2470-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#102470-prerelease",
		RuntimeVersion: "124.0.2470.0",
		Notes: `


<!-- ------------------------------ -->
#### Experimental APIs

The following Experimental APIs have been added in this Prerelease SDK.


<!-- ---------- -->
* Support for the Fluent Style Overlay Scrollbar.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ScrollBarStyle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.scrollbarstyle?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)

* [CoreWebView2ScrollbarStyle Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2scrollbarstyle?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * 'Default'
   * 'FluentOverlay'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ScrollBarStyle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#scrollbarstyle)

* [CoreWebView2ScrollbarStyle Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2scrollbarstyle?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true)
   * 'Default'
   * 'FluentOverlay'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalEnvironmentOptions2](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.2470-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalEnvironmentOptions2::get_ScrollBarStyle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.2470-prerelease&preserve-view=true#get_scrollbarstyle)
   * [ICoreWebView2ExperimentalEnvironmentOptions2::put_ScrollBarStyle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.2470-prerelease&preserve-view=true#put_scrollbarstyle)

* [COREWEBVIEW2_SCROLLBAR_STYLE enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.2470-prerelease&preserve-view=true#corewebview2_scrollbar_style)
   * 'COREWEBVIEW2_SCROLLBAR_STYLE_DEFAULT'
   * 'COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY'

---


<!-- ---------- -->
* Updated the WebMessageObjects API to allow injecting DOM objects into WebView2 content that's constructed via the app and via the 'CoreWebView2.PostWebMessage' API in the other direction.  Added a new web object type to represent a file system handle that can be posted to the web content to provide it with filesystem access.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.PostWebMessageAsJsonWithAdditionalObjects Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.postwebmessageasjsonwithadditionalobjects?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)

* 'CoreWebView2Environment' Class:
   * [CoreWebView2Environment.CreateWebFileSystemDirectoryHandle Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createwebfilesystemdirectoryhandle?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * [CoreWebView2Environment.CreateWebFileSystemFileHandle Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createwebfilesystemfilehandle?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)

* 'CoreWebView2FileSystemHandle' Class:
   * [CoreWebView2FileSystemHandle.Kind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandle.kind?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * [CoreWebView2FileSystemHandle.Path Property](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandle.path?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * [CoreWebView2FileSystemHandle.Permission Property](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandle.permission?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)

* [CoreWebView2FileSystemHandleKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandlekind?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * 'File'
   * 'Directory'

* [CoreWebView2FileSystemHandlePermission Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2filesystemhandlepermission?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * 'ReadOnly'
   * 'ReadWrite'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.PostWebMessageAsJsonWithAdditionalObjects Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#postwebmessageasjsonwithadditionalobjects)

* 'CoreWebView2Environment' Class:
   * [CoreWebView2Environment.CreateWebFileSystemDirectoryHandle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#createwebfilesystemdirectoryhandle)
   * [CoreWebView2Environment.CreateWebFileSystemFileHandle Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#createwebfilesystemfilehandle)

* 'CoreWebView2FileSystemHandle' Class:
   * [CoreWebView2FileSystemHandle.Kind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandle?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#kind)
   * [CoreWebView2FileSystemHandle.Path Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandle?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#path)
   * [CoreWebView2FileSystemHandle.Permission Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandle?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#permission)

* [CoreWebView2FileSystemHandleKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandlekind?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true)
   * 'File'
   * 'Directory'

* [CoreWebView2FileSystemHandlePermission Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2filesystemhandlepermission?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true)
   * 'ReadOnly'
   * 'ReadWrite'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental24](/microsoft-edge/webview2/reference/win32/icorewebview2experimental24?view=webview2-1.0.2470-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental24::PostWebMessageAsJsonWithAdditionalObjects](/microsoft-edge/webview2/reference/win32/icorewebview2experimental24?view=webview2-1.0.2470-prerelease&preserve-view=true#postwebmessageasjsonwithadditionalobjects)

* [ICoreWebView2ExperimentalEnvironment14](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment14?view=webview2-1.0.2470-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalEnvironment14::CreateObjectCollection](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment14?view=webview2-1.0.2470-prerelease&preserve-view=true#createobjectcollection)<!--win32 only-->
   * [ICoreWebView2ExperimentalEnvironment14::CreateWebFileSystemDirectoryHandle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment14?view=webview2-1.0.2470-prerelease&preserve-view=true#createwebfilesystemdirectoryhandle)
   * [ICoreWebView2ExperimentalEnvironment14::CreateWebFileSystemFileHandle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment14?view=webview2-1.0.2470-prerelease&preserve-view=true#createwebfilesystemfilehandle)

* [ICoreWebView2ExperimentalFileSystemHandle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalfilesystemhandle?view=webview2-1.0.2470-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalFileSystemHandle::get_Kind](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalfilesystemhandle?view=webview2-1.0.2470-prerelease&preserve-view=true#get_kind)
   * [ICoreWebView2ExperimentalFileSystemHandle::get_Path](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalfilesystemhandle?view=webview2-1.0.2470-prerelease&preserve-view=true#get_path)
   * [ICoreWebView2ExperimentalFileSystemHandle::get_Permission](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalfilesystemhandle?view=webview2-1.0.2470-prerelease&preserve-view=true#get_permission)

* [ICoreWebView2ExperimentalObjectCollection](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalobjectcollection?view=webview2-1.0.2470-prerelease&preserve-view=true)<!--win32 only-->
   * [ICoreWebView2ExperimentalObjectCollection::InsertValueAtIndex](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalobjectcollection?view=webview2-1.0.2470-prerelease&preserve-view=true#insertvalueatindex)<!--win32 only-->
   * [ICoreWebView2ExperimentalObjectCollection::RemoveValueAtIndex](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalobjectcollection?view=webview2-1.0.2470-prerelease&preserve-view=true#removevalueatindex)<!--win32 only-->

* [COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.2470-prerelease&preserve-view=true#corewebview2_file_system_handle_kind)
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND_FILE'
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND_DIRECTORY'

* [COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.2470-prerelease&preserve-view=true#corewebview2_file_system_handle_permission)
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_ONLY'
   * 'COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_WRITE'

---


<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ---------- -->
* Added the Runtime selection feature to support more prerelease testing and flighting scenarios.  You can specify 'ReleaseChannels' to choose which channels are searched for during environment creation, and 'ChannelSearchKind' to select a search order.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ChannelSearchKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.channelsearchkind?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * [CoreWebView2EnvironmentOptions.ReleaseChannels Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.releasechannels?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)

* [CoreWebView2ChannelSearchKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2channelsearchkind?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * 'MostStable'
   * 'LeastStable'

* [CoreWebView2ReleaseChannels Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2releasechannels?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)
   * 'None'
   * 'Stable'
   * 'Beta'
   * 'Dev'
   * 'Canary'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.ChannelSearchKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#channelsearchkind)
   * [CoreWebView2EnvironmentOptions.ReleaseChannels Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#releasechannels)

* [CoreWebView2ChannelSearchKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2channelsearchkind?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true)
   * 'MostStable'
   * 'LeastStable'

* [CoreWebView2ReleaseChannels Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2releasechannels?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true)
   * 'None'
   * 'Stable'
   * 'Beta'
   * 'Dev'
   * 'Canary'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions7](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2470-prerelease&preserve-view=true)
   * [ICoreWebView2EnvironmentOptions7::get_ChannelSearchKind](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2470-prerelease&preserve-view=true#get_channelsearchkind)
   * [ICoreWebView2EnvironmentOptions7::put_ChannelSearchKind](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2470-prerelease&preserve-view=true#put_channelsearchkind)
   * [ICoreWebView2EnvironmentOptions7::get_ReleaseChannels](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2470-prerelease&preserve-view=true#get_releasechannels)
   * [ICoreWebView2EnvironmentOptions7::put_ReleaseChannels](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7?view=webview2-1.0.2470-prerelease&preserve-view=true#put_releasechannels)

* [COREWEBVIEW2_CHANNEL_SEARCH_KIND enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2470-prerelease&preserve-view=true#corewebview2_channel_search_kind)
   * 'COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE'
   * 'COREWEBVIEW2_CHANNEL_SEARCH_KIND_LEAST_STABLE'

* [COREWEBVIEW2_RELEASE_CHANNELS enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2470-prerelease&preserve-view=true#corewebview2_release_channels)
   * 'COREWEBVIEW2_RELEASE_CHANNELS_NONE'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_STABLE'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_BETA'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_DEV'
   * 'COREWEBVIEW2_RELEASE_CHANNELS_CANARY'

---


<!-- ---------- -->
* Added the 'FailureSourceModulePath' property to the 'ProcessFailedEventArgs' type, to specify the full path of the module that caused the crash in cases of Windows code integrity failures - that is, when a process exited with 'STATUS_INVALID_IMAGE_HASH'.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2ProcessFailedEventArgs' Class:
   * [CoreWebView2ProcessFailedEventArgs.FailureSourceModulePath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2processfailedeventargs.failuresourcemodulepath?view=webview2-dotnet-1.0.2470-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2ProcessFailedEventArgs' Class:
   * [CoreWebView2ProcessFailedEventArgs.FailureSourceModulePath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2processfailedeventargs?view=webview2-winrt-1.0.2470-prerelease&preserve-view=true#failuresourcemodulepath)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ProcessFailedEventArgs3](/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs3?view=webview2-1.0.2470-prerelease&preserve-view=true)
   * [ICoreWebView2ProcessFailedEventArgs3::get_FailureSourceModulePath](/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs3?view=webview2-1.0.2470-prerelease&preserve-view=true#get_failuresourcemodulepath)<!--no put-->

---


<!-- ------------------------------ -->
#### Bug fixes


###### Runtime-only

* Fixed a reliability regression that could crash the application process when an old version of WebView2 client DLL is unloaded.
* Ensured that the WebView2 temporary download folder is unique per user data folder, and doesn't interfere with other apps or the browser.

<!-- end of Mar. 2024 Prerelease SDK -->


<!-- ====================================================================== -->
Release Date: February 26, 2024`,
	},
	"1.0.2365.46": {
		SDKVersion:     "1.0.2365.46",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10236546",
		RuntimeVersion: "122.0.2365.46",
		Notes: `

<!-- ------------------------------ -->
#### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
*  Added support for 'WebResourceRequested' for workers, which allows setting filters in order to receive 'WebResourceRequested' events for service workers, shared workers, and different-origin iframes.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.AddWebResourceRequestedFilter Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.addwebresourcerequestedfilter?view=webview2-dotnet-1.0.2365.46&preserve-view=true)
   * [CoreWebView2.RemoveWebResourceRequestedFilter Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.removewebresourcerequestedfilter?view=webview2-dotnet-1.0.2365.46&preserve-view=true)

* 'CoreWebView2WebResourceRequestedEventArgs' Class:
   * [CoreWebView2WebResourceRequestedEventArgs.RequestedSourceKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2webresourcerequestedeventargs.requestedsourcekind?view=webview2-dotnet-1.0.2365.46&preserve-view=true)

* [CoreWebView2WebResourceRequestSourceKinds Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2webresourcerequestsourcekinds?view=webview2-dotnet-1.0.2365.46&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.AddWebResourceRequestedFilter Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2365.46&preserve-view=true#addwebresourcerequestedfilter)
   * [CoreWebView2.RemoveWebResourceRequestedFilter Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.2365.46&preserve-view=true#removewebresourcerequestedfilter)

* 'CoreWebView2WebResourceRequestedEventArgs' Class:
   * [CoreWebView2WebResourceRequestedEventArgs.RequestedSourceKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webresourcerequestedeventargs?view=webview2-winrt-1.0.2365.46&preserve-view=true#requestedsourcekind)

* [CoreWebView2WebResourceRequestSourceKinds Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webresourcerequestsourcekinds?view=webview2-winrt-1.0.2365.46&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_22](/microsoft-edge/webview2/reference/win32/icorewebview2_22?view=webview2-1.0.2365.46&preserve-view=true)
    * [ICoreWebView2_22::AddWebResourceRequestedFilterWithRequestSourceKinds](/microsoft-edge/webview2/reference/win32/icorewebview2_22?view=webview2-1.0.2365.46&preserve-view=true#addwebresourcerequestedfilterwithrequestsourcekinds)
    * [ICoreWebView2_22::RemoveWebResourceRequestedFilterWithRequestSourceKinds](/microsoft-edge/webview2/reference/win32/icorewebview2_22?view=webview2-1.0.2365.46&preserve-view=true#removewebresourcerequestedfilterwithrequestsourcekinds)

* [ICoreWebView2WebResourceRequestedEventArgs2](/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs2?view=webview2-1.0.2365.46&preserve-view=true)
    * [ICoreWebView2WebResourceRequestedEventArgs2::get_RequestedSourceKind](/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs2?view=webview2-1.0.2365.46&preserve-view=true#get_requestedsourcekind)

* [COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.2365.46&preserve-view=true#corewebview2_web_resource_request_source_kinds)

---


<!-- ------------------------------ -->
* To support browser extensions in WebView2, added 'GetBrowserExtensions' for WinRT:

##### [.NET/C#](#tab/dotnetcsharp)

N/A

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2Profile' Class:
    * [CoreWebView2Profile.GetBrowserExtensionsAsync Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.2365.46&preserve-view=true#getbrowserextensionsasync)

##### [Win32/C++](#tab/win32cpp)

N/A

---


<!-- ------------------------------ -->
#### Bug fixes

###### Runtime-only

* Fixed a regression that affected handling of the 'NewWindowRequested' event when the new window is set to be the source WebView.  ([Issue #4250](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4250))

* Fixed a bug where closing a WebView that has an embedded PDF viewer could lead to a crash.  ([Issue #3832](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3832))

* Fixed a regression where mouse-clicks stopped working when the application enabled 'SetWindowDisplayAffinity'.  ([Issue #4325](https://github.com/MicrosoftEdge/WebView2Feedback/issues/4325))

<!-- end of Feb 2024 Release SDK -->


<!-- ====================================================================== -->
Release Date: February 26, 2024`,
	},
}
