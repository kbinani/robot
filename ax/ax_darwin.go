// Package ax is a wrapper of accessibility API.
package ax

// #cgo LDFLAGS: -framework ApplicationServices
// #include <ApplicationServices/ApplicationServices.h>
import "C"

import (
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"unicode/utf16"
	"unsafe"
)

const (
	// Attributes
	AMPMFieldAttribute                  = "AXAMPMField"                  // AXUIElement
	AllowedValuesAttribute              = "AXAllowedValues"              // CFArray<[same as kAXValueAttribute]>
	AlternateUIVisibleAttribute         = "AXAlternateUIVisible"         // CFBoolean
	CancelButtonAttribute               = "AXCancelButton"               // AXUIElement
	ChildrenAttribute                   = "AXChildren"                   // CFArray<AXUIElement>
	ClearButtonAttribute                = "AXClearButton"                // AXUIElement
	CloseButtonAttribute                = "AXCloseButton"                // AXUIElement
	ColumnCountAttribute                = "AXColumnCount"                // CFNumber<SInt64>
	ColumnHeaderUIElementsAttribute     = "AXColumnHeaderUIElements"     // CFArray<AXUIElement>
	ColumnIndexRangeAttribute           = "AXColumnIndexRange"           // AXValue<CFRange>
	ColumnTitleAttribute                = "AXColumnTitles"               // CFArray<CFString>
	ColumnTitlesAttribute               = "AXColumnTitles"               // CFArray<CFString>
	ColumnsAttribute                    = "AXColumns"                    // CFArray<AXUIElement>
	ContentsAttribute                   = "AXContents"                   // CFArray<AXUIElement>
	CriticalValueAttribute              = "AXCriticalValue"              // [same as kAXValueAttribute]
	DayFieldAttribute                   = "AXDayField"                   // AXUIElement
	DecrementButtonAttribute            = "AXDecrementButton"            // AXUIElement
	DefaultButtonAttribute              = "AXDefaultButton"              // AXUIElement
	DescriptionAttribute                = "AXDescription"                // CFString
	DisclosedByRowAttribute             = "AXDisclosedByRow"             // AXUIElement
	DisclosedRowsAttribute              = "AXDisclosedRows"              // CFArray<AXUIElement>
	DisclosingAttribute                 = "AXDisclosing"                 // CFBoolean
	DisclosureLevelAttribute            = "AXDisclosureLevel"            // CFNumber<SInt64>
	DocumentAttribute                   = "AXDocument"                   // CFString
	EditedAttribute                     = "AXEdited"                     // CFBoolean
	ElementBusyAttribute                = "AXElementBusy"                // CFBoolean
	EnabledAttribute                    = "AXEnabled"                    // CFBoolean
	ExpandedAttribute                   = "AXExpanded"                   // CFBoolean
	ExtrasMenuBarAttribute              = "AXExtrasMenuBar"              // AXUIElement
	FilenameAttribute                   = "AXFilename"                   // CFString
	FocusedApplicationAttribute         = "AXFocusedApplication"         // AXUIElement
	FocusedAttribute                    = "AXFocused"                    // CFBoolean
	FocusedUIElementAttribute           = "AXFocusedUIElement"           // AXUIElement
	FocusedWindowAttribute              = "AXFocusedWindow"              // AXUIElement
	FrontmostAttribute                  = "AXFrontmost"                  // CFBoolean
	FullScreenButtonAttribute           = "AXFullScreenButton"           // AXUIElement
	GrowAreaAttribute                   = "AXGrowArea"                   // AXUIElement
	HandlesAttribute                    = "AXHandles"                    // CFArray<AXUIElement>
	HeaderAttribute                     = "AXHeader"                     // AXUIElement
	HelpAttribute                       = "AXHelp"                       // CFString
	HiddenAttribute                     = "AXHidden"                     // CFBoolean
	HorizontalScrollBarAttribute        = "AXHorizontalScrollBar"        // AXUIElement
	HorizontalUnitDescriptionAttribute  = "AXHorizontalUnitDescription"  // CFString
	HorizontalUnitsAttribute            = "AXHorizontalUnits"            // CFString
	HourFieldAttribute                  = "AXHourField"                  // AXUIElement
	IdentifierAttribute                 = "AXIdentifier"                 // CFString
	IncrementButtonAttribute            = "AXIncrementButton"            // AXUIElement
	IncrementorAttribute                = "AXIncrementor"                // AXUIElement
	IndexAttribute                      = "AXIndex"                      // CFNumber<SInt64>
	InsertionPointLineNumberAttribute   = "AXInsertionPointLineNumber"   // CFNumber<SInt64>
	IsApplicationRunningAttribute       = "AXIsApplicationRunning"       // CFBoolean
	IsEditableAttribute                 = "AXIsEditable"                 // CFBoolean
	LabelUIElementsAttribute            = "AXLabelUIElements"            // CFArray<AXUIElement>
	LabelValueAttribute                 = "AXLabelValue"                 // CFNumber<Float32>
	LinkedUIElementsAttribute           = "AXLinkedUIElements"           // CFArray<AXUIElement>
	MainAttribute                       = "AXMain"                       // CFBoolean
	MainWindowAttribute                 = "AXMainWindow"                 // AXUIElement
	MarkerTypeAttribute                 = "AXMarkerType"                 // CFString
	MarkerTypeDescriptionAttribute      = "AXMarkerTypeDescription"      // CFString
	MarkerUIElementsAttribute           = "AXMarkerUIElements"           // CFArray<AXUIElement>
	MatteContentUIElementAttribute      = "AXMatteContentUIElement"      // AXUIElement
	MatteHoleAttribute                  = "AXMatteHole"                  // AXValue<CGRect>
	MaxValueAttribute                   = "AXMaxValue"                   // [same as kAXValueAttribute]
	MenuBarAttribute                    = "AXMenuBar"                    // AXUIElement
	MenuItemCmdCharAttribute            = "AXMenuItemCmdChar"            // CFString
	MenuItemCmdGlyphAttribute           = "AXMenuItemCmdGlyph"           // CFNumber<SInt64>
	MenuItemCmdModifiersAttribute       = "AXMenuItemCmdModifiers"       // CFNumber<SInt64>
	MenuItemCmdVirtualKeyAttribute      = "AXMenuItemCmdVirtualKey"      // CFNumber<SInt64>
	MenuItemMarkCharAttribute           = "AXMenuItemMarkChar"           // CFString
	MenuItemPrimaryUIElementAttribute   = "AXMenuItemPrimaryUIElement"   // AXUIElement
	MinValueAttribute                   = "AXMinValue"                   // [same as kAXValueAttribute]
	MinimizeButtonAttribute             = "AXMinimizeButton"             // AXUIElement
	MinimizedAttribute                  = "AXMinimized"                  // CFBoolean
	MinuteFieldAttribute                = "AXMinuteField"                // AXUIElement
	ModalAttribute                      = "AXModal"                      // CFBoolean
	MonthFieldAttribute                 = "AXMonthField"                 // AXUIElement
	NextContentsAttribute               = "AXNextContents"               // CFArray<AXUIElement>
	NumberOfCharactersAttribute         = "AXNumberOfCharacters"         // CFNumber<SInt64>
	OrderedByRowAttribute               = "AXOrderedByRow"               // CFBoolean
	OrientationAttribute                = "AXOrientation"                // CFString
	OverflowButtonAttribute             = "AXOverflowButton"             // AXUIElement
	ParentAttribute                     = "AXParent"                     // AXUIElement
	PlaceholderValueAttribute           = "AXPlaceholderValue"           // CFString
	PositionAttribute                   = "AXPosition"                   // AXValue<CGPoint>
	PreviousContentsAttribute           = "AXPreviousContents"           // CFArray<AXUIElement>
	ProxyAttribute                      = "AXProxy"                      // AXUIElement
	RoleAttribute                       = "AXRole"                       // CFString
	RoleDescriptionAttribute            = "AXRoleDescription"            // CFString
	RowCountAttribute                   = "AXRowCount"                   // CFNumber<SInt64>
	RowHeaderUIElementsAttribute        = "AXRowHeaderUIElements"        // CFArray<AXUIElement>
	RowIndexRangeAttribute              = "AXRowIndexRange"              // AXValue<CFRange>
	RowsAttribute                       = "AXRows"                       // CFArray<AXUIElement>
	SearchButtonAttribute               = "AXSearchButton"               // AXUIElement
	SecondFieldAttribute                = "AXSecondField"                // AXUIElement
	SelectedAttribute                   = "AXSelected"                   // CFBoolean
	SelectedCellsAttribute              = "AXSelectedCells"              // CFArray<AXUIElement>
	SelectedChildrenAttribute           = "AXSelectedChildren"           // CFArray<AXUIElement>
	SelectedColumnsAttribute            = "AXSelectedColumns"            // CFArray<AXUIElement>
	SelectedRowsAttribute               = "AXSelectedRows"               // CFArray<AXUIElement>
	SelectedTextAttribute               = "AXSelectedText"               // CFString
	SelectedTextRangeAttribute          = "AXSelectedTextRange"          // AXValue<CFRange>
	SelectedTextRangesAttribute         = "AXSelectedTextRanges"         // CFArray<AXValue<CFRange>>
	ServesAsTitleForUIElementsAttribute = "AXServesAsTitleForUIElements" // CFArray<AXUIElement>
	SharedCharacterRangeAttribute       = "AXSharedCharacterRange"       // AXValue<CFRange>
	SharedFocusElementsAttribute        = "AXSharedFocusElements"        // CFArray<AXUIElement>
	SharedTextUIElementsAttribute       = "AXSharedTextUIElements"       // CFArray<AXUIElement>
	ShownMenuUIElementAttribute         = "AXShownMenuUIElement"         // AXUIElement
	SizeAttribute                       = "AXSize"                       // AXValue<CGSize>
	SortDirectionAttribute              = "AXSortDirection"              // CFString
	SplittersAttribute                  = "AXSplitters"                  // CFArray<AXUIElement>
	SubroleAttribute                    = "AXSubrole"                    // CFString
	TabsAttribute                       = "AXTabs"                       // CFArray<AXUIElement>
	TextAttribute                       = "AXText"                       // CFString
	TitleAttribute                      = "AXTitle"                      // CFString
	TitleUIElementAttribute             = "AXTitleUIElement"             // AXUIElement
	ToolbarButtonAttribute              = "AXToolbarButton"              // AXUIElement
	TopLevelUIElementAttribute          = "AXTopLevelUIElement"          // AXUIElement
	URLAttribute                        = "AXURL"                        // CFURL
	UnitDescriptionAttribute            = "AXUnitDescription"            // CFString
	UnitsAttribute                      = "AXUnits"                      // CFString
	ValueAttribute                      = "AXValue"                      // [same as kAXValueAttribute]
	ValueDescriptionAttribute           = "AXValueDescription"           // CFString
	ValueIncrementAttribute             = "AXValueIncrement"             // [same as kAXValueAttribute]
	ValueWrapsAttribute                 = "AXValueWraps"                 // NSNumber with BOOL value
	VerticalScrollBarAttribute          = "AXVerticalScrollBar"          // AXUIElement
	VerticalUnitDescriptionAttribute    = "AXVerticalUnitDescription"    // CFString
	VerticalUnitsAttribute              = "AXVerticalUnits"              // CFString
	VisibleCellsAttribute               = "AXVisibleCells"               // CFArray<AXUIElement>
	VisibleCharacterRangeAttribute      = "AXVisibleCharacterRange"      // AXValue<CFRange>
	VisibleChildrenAttribute            = "AXVisibleChildren"            // CFArray<AXUIElement>
	VisibleColumnsAttribute             = "AXVisibleColumns"             // CFArray<AXUIElement>
	VisibleRowsAttribute                = "AXVisibleRows"                // CFArray<AXUIElement>
	VisibleTextAttribute                = "AXVisibleText"                // CFString
	WarningValueAttribute               = "AXWarningValue"               // [same as kAXValueAttribute]
	WindowAttribute                     = "AXWindow"                     // AXUIElement
	WindowsAttribute                    = "AXWindows"                    // CFArray<AXUIElement>
	YearFieldAttribute                  = "AXYearField"                  // AXUIElement
	ZoomButtonAttribute                 = "AXZoomButton"                 // AXUIElement

	// Parameterized Attributes
	AttributedStringForRangeParameterizedAttribute  = "AXAttributedStringForRange"  // CFAttributedString
	BoundsForRangeParameterizedAttribute            = "AXBoundsForRange"            // AXValue<CGRect>
	CellForColumnAndRowParameterizedAttribute       = "AXCellForColumnAndRow"       // AXUIElement
	LayoutPointForScreenPointParameterizedAttribute = "AXLayoutPointForScreenPoint" // AXValue<CGPoint>
	LayoutSizeForScreenSizeParameterizedAttribute   = "AXLayoutSizeForScreenSize"   // AXValue<CGSize>
	LineForIndexParameterizedAttribute              = "AXLineForIndex"              // CFNumber<SInt64>
	RTFForRangeParameterizedAttribute               = "AXRTFForRange"               // CFData
	RangeForIndexParameterizedAttribute             = "AXRangeForIndex"             // AXValue<CFRange>
	RangeForLineParameterizedAttribute              = "AXRangeForLine"              // AXValue<CFRange>
	RangeForPositionParameterizedAttribute          = "AXRangeForPosition"          // AXValue<CFRange>
	ScreenPointForLayoutPointParameterizedAttribute = "AXScreenPointForLayoutPoint" // AXValue<CGPoint>
	ScreenSizeForLayoutSizeParameterizedAttribute   = "AXScreenSizeForLayoutSize"   // AXValue<CGSize>
	StringForRangeParameterizedAttribute            = "AXStringForRange"            // CFString
	StyleRangeForIndexParameterizedAttribute        = "AXStyleRangeForIndex"        // AXValue<CFRange>

	// Actions
	PressAction           = "AXPress"
	IncrementAction       = "AXIncrement"
	DecrementAction       = "AXDecrement"
	ConfirmAction         = "AXConfirm"
	CancelAction          = "AXCancel"
	ShowAlternateUIAction = "AXShowAlternateUI"
	ShowDefaultUIAction   = "AXShowDefaultUI"
	RaiseAction           = "AXRaise"
	ShowMenuAction        = "AXShowMenu"
	PickAction            = "AXPick"

	// Roles
	ApplicationRole        = "AXApplication"
	SystemWideRole         = "AXSystemWide"
	WindowRole             = "AXWindow"
	SheetRole              = "AXSheet"
	DrawerRole             = "AXDrawer"
	GrowAreaRole           = "AXGrowArea"
	ImageRole              = "AXImage"
	UnknownRole            = "AXUnknown"
	ButtonRole             = "AXButton"
	RadioButtonRole        = "AXRadioButton"
	CheckBoxRole           = "AXCheckBox"
	PopUpButtonRole        = "AXPopUpButton"
	MenuButtonRole         = "AXMenuButton"
	TabGroupRole           = "AXTabGroup"
	TableRole              = "AXTable"
	ColumnRole             = "AXColumn"
	RowRole                = "AXRow"
	OutlineRole            = "AXOutline"
	BrowserRole            = "AXBrowser"
	ScrollAreaRole         = "AXScrollArea"
	ScrollBarRole          = "AXScrollBar"
	RadioGroupRole         = "AXRadioGroup"
	ListRole               = "AXList"
	GroupRole              = "AXGroup"
	ValueIndicatorRole     = "AXValueIndicator"
	ComboBoxRole           = "AXComboBox"
	SliderRole             = "AXSlider"
	IncrementorRole        = "AXIncrementor"
	BusyIndicatorRole      = "AXBusyIndicator"
	ProgressIndicatorRole  = "AXProgressIndicator"
	RelevanceIndicatorRole = "AXRelevanceIndicator"
	ToolbarRole            = "AXToolbar"
	DisclosureTriangleRole = "AXDisclosureTriangle"
	TextFieldRole          = "AXTextField"
	TextAreaRole           = "AXTextArea"
	StaticTextRole         = "AXStaticText"
	MenuBarRole            = "AXMenuBar"
	MenuBarItemRole        = "AXMenuBarItem"
	MenuRole               = "AXMenu"
	MenuItemRole           = "AXMenuItem"
	SplitGroupRole         = "AXSplitGroup"
	SplitterRole           = "AXSplitter"
	ColorWellRole          = "AXColorWell"
	TimeFieldRole          = "AXTimeField"
	DateFieldRole          = "AXDateField"
	HelpTagRole            = "AXHelpTag"
	MatteRole              = "AXMatte"
	DockItemRole           = "AXDockItem"
	RulerRole              = "AXRuler"
	RulerMarkerRole        = "AXRulerMarker"
	GridRole               = "AXGrid"
	LevelIndicatorRole     = "AXLevelIndicator"
	CellRole               = "AXCell"
	LayoutAreaRole         = "AXLayoutArea"
	LayoutItemRole         = "AXLayoutItem"
	HandleRole             = "AXHandle"
	PopoverRole            = "AXPopover"

	// Subroles
	CloseButtonSubrole             = "AXCloseButton"
	MinimizeButtonSubrole          = "AXMinimizeButton"
	ZoomButtonSubrole              = "AXZoomButton"
	ToolbarButtonSubrole           = "AXToolbarButton"
	FullScreenButtonSubrole        = "AXFullScreenButton"
	SecureTextFieldSubrole         = "AXSecureTextField"
	TableRowSubrole                = "AXTableRow"
	OutlineRowSubrole              = "AXOutlineRow"
	UnknownSubrole                 = "AXUnknown"
	StandardWindowSubrole          = "AXStandardWindow"
	DialogSubrole                  = "AXDialog"
	SystemDialogSubrole            = "AXSystemDialog"
	FloatingWindowSubrole          = "AXFloatingWindow"
	SystemFloatingWindowSubrole    = "AXSystemFloatingWindow"
	IncrementArrowSubrole          = "AXIncrementArrow"
	DecrementArrowSubrole          = "AXDecrementArrow"
	IncrementPageSubrole           = "AXIncrementPage"
	DecrementPageSubrole           = "AXDecrementPage"
	SortButtonSubrole              = "AXSortButton"
	SearchFieldSubrole             = "AXSearchField"
	TimelineSubrole                = "AXTimeline"
	RatingIndicatorSubrole         = "AXRatingIndicator"
	ContentListSubrole             = "AXContentList"
	DefinitionListSubrole          = "AXDefinitionList"
	DescriptionListSubrole         = "AXDescriptionList"
	ToggleSubrole                  = "AXToggle"
	SwitchSubrole                  = "AXSwitch"
	ApplicationDockItemSubrole     = "AXApplicationDockItem"
	DocumentDockItemSubrole        = "AXDocumentDockItem"
	FolderDockItemSubrole          = "AXFolderDockItem"
	MinimizedWindowDockItemSubrole = "AXMinimizedWindowDockItem"
	URLDockItemSubrole             = "AXURLDockItem"
	DockExtraDockItemSubrole       = "AXDockExtraDockItem"
	TrashDockItemSubrole           = "AXTrashDockItem"
	SeparatorDockItemSubrole       = "AXSeparatorDockItem"
	ProcessSwitcherListSubrole     = "AXProcessSwitcherList"
)

// UIElement is a container of AXUIElementRef.
type UIElement struct {
	obj C.AXUIElementRef
}

// Range is a wrapper for CFRange.
type Range struct {
	Location int64
	Length   int64
}

// Point is a wrapper for CGPoint.
type Point struct {
	X float64
	Y float64
}

// Size is a wrapper for CGSize.
type Size struct {
	Dx float64
	Dy float64
}

// Rect is a wrapper for CGRect.
type Rect struct {
	Origin Point
	Size   Size
}

// Actions returns available actions for AXUIElementRef.
func (ref *UIElement) Actions() []string {
	a := []string{}
	var actions C.CFArrayRef
	C.AXUIElementCopyActionNames(ref.obj, &actions)
	if actions == nil {
		return a
	}
	defer C.CFRelease((C.CFTypeRef)(actions))
	num := int(C.CFArrayGetCount(actions))
	for i := 0; i < num; i++ {
		item := C.CFArrayGetValueAtIndex(actions, C.CFIndex(i))
		a = append(a, stringFromCFString((C.CFStringRef)(item)))
	}
	return a
}

// CreateApplication create UIElement by process id.
func CreateApplication(pid int) *UIElement {
	ref := C.AXUIElementCreateApplication(C.pid_t(pid))
	return newUIElement(ref)
}

// CreateSystemWide create system-wide of UIElement.
func CreateSystemWide() *UIElement {
	ref := C.AXUIElementCreateSystemWide()
	return newUIElement(ref)
}

// Perform is a wrapper func for AXUIElementPerformAction.
func (ref *UIElement) Perform(action string) {
	if ref == nil {
		return
	}
	if ref.obj == nil {
		return
	}
	a := cfstr(action)
	defer C.CFRelease((C.CFTypeRef)(a))
	C.AXUIElementPerformAction(ref.obj, a)
}

func convertCFType(obj C.CFTypeRef) interface{} {
	if obj == nil {
		return nil
	}

	id := C.CFGetTypeID(obj)
	switch id {
	case C.CFStringGetTypeID():
		s := stringFromCFString(C.CFStringRef(obj))
		C.CFRelease(obj)
		return s
	case C.CFBooleanGetTypeID():
		b := C.CFBooleanRef(obj) == C.kCFBooleanTrue
		C.CFRelease(obj)
		return b
	case C.CFNumberGetTypeID():
		t := C.CFNumberGetType(C.CFNumberRef(obj))
		defer C.CFRelease(obj)

		switch t {
		case C.kCFNumberCharType:
			var v C.char
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return byte(v)
		case C.kCFNumberSInt8Type:
			var v int8
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return v
		case C.kCFNumberSInt16Type:
			var v int16
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return v
		case C.kCFNumberSInt32Type:
			var v int32
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return v
		case C.kCFNumberSInt64Type:
			var v int64
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return v
		case C.kCFNumberFloat32Type, C.kCFNumberFloatType:
			var v float32
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return v
		case C.kCFNumberFloat64Type, C.kCFNumberDoubleType:
			var v float64
			C.CFNumberGetValue(C.CFNumberRef(obj), t, unsafe.Pointer(&v))
			return v
		default:
			return nil
			// case C.kCFNumberShortType:
			// case C.kCFNumberIntType:
			// case C.kCFNumberLongType:
			// case C.kCFNumberLongLongType:
			// case C.kCFNumberCFIndexType:
			// case C.kCFNumberNSIntegerType:
			// case C.kCFNumberCGFloatType:
		}
	case C.AXValueGetTypeID():
		defer C.CFRelease(obj)
		t := C.AXValueGetType(C.AXValueRef(obj))
		switch t {
		case C.kAXValueTypeCGPoint:
			var v C.CGPoint
			if C.AXValueGetValue(C.AXValueRef(obj), t, unsafe.Pointer(&v)) == 0 {
				return nil
			}
			return Point{float64(v.x), float64(v.y)}
		case C.kAXValueTypeCGSize:
			var v C.CGSize
			if C.AXValueGetValue(C.AXValueRef(obj), t, unsafe.Pointer(&v)) == 0 {
				return nil
			}
			return Size{float64(v.width), float64(v.height)}
		case C.kAXValueTypeCGRect:
			var v C.CGRect
			if C.AXValueGetValue(C.AXValueRef(obj), t, unsafe.Pointer(&v)) == 0 {
				return nil
			}
			return Rect{Point{float64(v.origin.x), float64(v.origin.y)}, Size{float64(v.size.width), float64(v.size.height)}}
		case C.kAXValueTypeCFRange:
			var v C.CFRange
			if C.AXValueGetValue(C.AXValueRef(obj), t, unsafe.Pointer(&v)) == 0 {
				return nil
			}
			return Range{int64(v.location), int64(v.length)}
		default:
			return nil
		}
	case C.CFURLGetTypeID():
		defer C.CFRelease(obj)
		path := C.CFURLGetString(C.CFURLRef(obj))
		p := stringFromCFString(path)
		u, _ := url.Parse(p)
		return u
	case C.AXUIElementGetTypeID():
		return newUIElement(C.AXUIElementRef(obj))
	}
	return nil
}

func (ref *UIElement) attr(attribute string, expectedType C.CFTypeID) (C.CFTypeRef, error) {
	a := cfstr(attribute)
	defer C.CFRelease(C.CFTypeRef(a))

	var value C.CFTypeRef
	C.AXUIElementCopyAttributeValue(ref.obj, a, &value)
	if value == nil {
		return nil, errors.New("AXUIElementCopyAttributeValue returns nil")
	}
	t := C.CFGetTypeID(value)
	if t != expectedType {
		C.CFRelease(value)
		return nil, errors.New("Return type mismatch")
	}

	return value, nil
}

func (ref *UIElement) SliceOfUIElementAttr(attribute string) []*UIElement {
	items := []*UIElement{}
	ret, err := ref.attr(attribute, C.CFArrayGetTypeID())
	if err != nil {
		return items
	}
	num := int(C.CFArrayGetCount(C.CFArrayRef(ret)))
	for i := 0; i < num; i++ {
		o := C.CFTypeRef(C.CFArrayGetValueAtIndex(C.CFArrayRef(ret), C.CFIndex(i)))
		if o == nil {
			continue
		}
		C.CFRetain(o)
		t := convertCFType(o)
		e, ok := t.(*UIElement)
		if !ok {
			continue
		}
		items = append(items, e)
	}
	return items
}

func (ref *UIElement) SliceOfRangeAttr(attribute string) []Range {
	items := []Range{}
	ret, err := ref.attr(attribute, C.CFArrayGetTypeID())
	if err != nil {
		return items
	}
	num := int(C.CFArrayGetCount(C.CFArrayRef(ret)))
	for i := 0; i < num; i++ {
		o := C.CFTypeRef(C.CFArrayGetValueAtIndex(C.CFArrayRef(ret), C.CFIndex(i)))
		if o == nil {
			continue
		}
		C.CFRetain(o)
		t := convertCFType(o)
		e, ok := t.(Range)
		if !ok {
			continue
		}
		items = append(items, e)
	}

	return items
}

func (ref *UIElement) SliceOfStringAttr(attribute string) []string {
	items := []string{}
	ret, err := ref.attr(attribute, C.CFArrayGetTypeID())
	if err != nil {
		return items
	}
	num := int(C.CFArrayGetCount(C.CFArrayRef(ret)))
	for i := 0; i < num; i++ {
		o := C.CFTypeRef(C.CFArrayGetValueAtIndex(C.CFArrayRef(ret), C.CFIndex(i)))
		if o == nil {
			continue
		}
		C.CFRetain(o)
		t := convertCFType(o)
		e, ok := t.(string)
		if !ok {
			continue
		}
		items = append(items, e)
	}
	return items
}

func (ref *UIElement) UIElementAttr(attribute string) (*UIElement, error) {
	ret, err := ref.attr(attribute, C.AXUIElementGetTypeID())
	if err != nil {
		return nil, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(*UIElement)
	if !ok {
		return nil, errors.New("Cannot convert CFTypeRef to *UIElement")
	}
	return t, nil
}

func (ref *UIElement) BoolAttr(attribute string) (bool, error) {
	ret, err := ref.attr(attribute, C.CFBooleanGetTypeID())
	var d bool = false
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(bool)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to bool")
	}
	return t, nil
}

func (ref *UIElement) StringAttr(attribute string) (string, error) {
	ret, err := ref.attr(attribute, C.CFStringGetTypeID())
	var d string = ""
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(string)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to string")
	}
	return t, nil
}

func (ref *UIElement) Int32Attr(attribute string) (int32, error) {
	ret, err := ref.attr(attribute, C.CFNumberGetTypeID())
	var d int32 = 0
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(int32)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to int32")
	}
	return t, nil
}

func (ref *UIElement) Int64Attr(attribute string) (int64, error) {
	ret, err := ref.attr(attribute, C.CFNumberGetTypeID())
	var d int64 = 0
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(int64)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to int64")
	}
	return t, nil
}

func (ref *UIElement) Float32Attr(attribute string) (float32, error) {
	ret, err := ref.attr(attribute, C.CFNumberGetTypeID())
	var d float32 = 0
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(float32)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to float32")
	}
	return t, nil
}

func (ref *UIElement) Float64Attr(attribute string) (float64, error) {
	ret, err := ref.attr(attribute, C.CFNumberGetTypeID())
	var d float64 = 0
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(float64)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to float64")
	}
	return t, nil
}

func (ref *UIElement) RangeAttr(attribute string) (Range, error) {
	ret, err := ref.attr(attribute, C.AXValueGetTypeID())
	var d Range = Range{0, 0}
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(Range)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to Range")
	}
	return t, nil
}

func (ref *UIElement) PointAttr(attribute string) (Point, error) {
	ret, err := ref.attr(attribute, C.AXValueGetTypeID())
	var d Point = Point{0, 0}
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(Point)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to Point")
	}
	return t, nil
}

func (ref *UIElement) SizeAttr(attribute string) (Size, error) {
	ret, err := ref.attr(attribute, C.AXValueGetTypeID())
	var d Size = Size{0, 0}
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(Size)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to Size")
	}
	return t, nil
}

func (ref *UIElement) RectAttr(attribute string) (Rect, error) {
	ret, err := ref.attr(attribute, C.AXValueGetTypeID())
	var d Rect = Rect{Point{0, 0}, Size{0, 0}}
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(Rect)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to Rect")
	}
	return t, nil
}

func (ref *UIElement) URLAttr(attribute string) (*url.URL, error) {
	ret, err := ref.attr(attribute, C.CFURLGetTypeID())
	var d *url.URL = nil
	if err != nil {
		return d, err
	}
	obj := convertCFType(ret)
	t, ok := obj.(*url.URL)
	if !ok {
		return d, errors.New("Cannot convert CFTypeRef to *url.URL")
	}
	return t, nil
}

func finalizeUIElement(ref *UIElement) {
	if ref != nil {
		ref.CFRelease()
	}
}

func newUIElement(obj C.AXUIElementRef) *UIElement {
	o := new(UIElement)
	o.obj = obj
	runtime.SetFinalizer(o, finalizeUIElement)
	return o
}

// CFRelease disposes AXUIElementRef object.
func (ref *UIElement) CFRelease() {
	if ref.obj != nil {
		C.CFRelease(ref.obj)
		ref.obj = nil
	}
}

func cfstr(s string) C.CFStringRef {
	ptr := C.CString(s)
	ret := C.CFStringCreateWithCString(nil, ptr, C.kCFStringEncodingUTF8)
	C.free(unsafe.Pointer(ptr))
	return ret
}

func stringFromCFString(s C.CFStringRef) string {
	ptr := C.CFStringGetCStringPtr(s, C.kCFStringEncodingUTF8)
	if ptr != nil {
		return C.GoString(ptr)
	}
	length := uint32(C.CFStringGetLength(s))
	uniPtr := C.CFStringGetCharactersPtr(s)
	if uniPtr == nil || length == 0 {
		return ""
	}
	return stringFromUnicode16Ptr((*uint16)(uniPtr), length)
}

func stringFromUnicode16Ptr(p *uint16, length uint32) string {
	r := []uint16{}
	ptr := uintptr(unsafe.Pointer(p))
	for i := uint32(0); i < length; i++ {
		c := *(*uint16)(unsafe.Pointer(ptr))
		r = append(r, c)
		if c == 0 {
			break
		}
		ptr = ptr + unsafe.Sizeof(c)
	}
	r = append(r, uint16(0))
	decoded := utf16.Decode(r)
	n := 0
	for i, r := range decoded {
		if r == rune(0) {
			n = i
			break
		}
	}
	return string(decoded[:n])
}

// ValueWraps returns kAXValueWrapsAttribute of AXUIElementRef.
func (ref *UIElement) ValueWraps() bool {
	ret, _ := ref.attr(ValueWrapsAttribute, C.CFNumberGetTypeID())
	var d bool = false
	obj := convertCFType(ret)
	if obj == nil {
		return d
	}
	t, ok := obj.(byte)
	if !ok {
		return d
	}
	return t != byte(0)
}

func (ref *UIElement) attrAny(attribute string) interface{} {
	a := cfstr(attribute)
	defer C.CFRelease(C.CFTypeRef(a))
	var value C.CFTypeRef
	C.AXUIElementCopyAttributeValue(ref.obj, a, &value)
	if value == nil {
		return nil
	}
	defer C.CFRelease(value)
	ret := convertCFType(value)
	return ret
}

// BoundsForRange returns kAXBoundsForRangeParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) BoundsForRange(r Range) Rect {
	a := cfstr(BoundsForRangeParameterizedAttribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfRange := C.CFRangeMake(C.CFIndex(r.Location), C.CFIndex(r.Length))
	axRangeValue := C.AXValueCreate(C.kAXValueTypeCFRange, unsafe.Pointer(&cfRange))
	defer C.CFRelease(C.CFTypeRef(axRangeValue))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, axRangeValue, &value)
	ret := convertCFType(value)
	o, ok := ret.(Rect)
	if !ok {
		return Rect{Point{0, 0}, Size{0, 0}}
	}
	return o
}

// CellForColumnAndRow returns kAXCellForColumnAndRowParameterizedAttribute of AXUIElement.
func (ref *UIElement) CellForColumnAndRow(column, row uint32) *UIElement {
	a := cfstr(CellForColumnAndRowParameterizedAttribute)
	defer C.CFRelease(C.CFTypeRef(a))

	params := C.CFArrayCreateMutable(nil, C.CFIndex(2), nil)
	defer C.CFRelease(C.CFTypeRef(params))

	var cColumn int64 = int64(column)
	cfColumn := C.CFNumberCreate(nil, C.kCFNumberSInt64Type, unsafe.Pointer(&cColumn))
	defer C.CFRelease(C.CFTypeRef(cfColumn))
	C.CFArrayAppendValue(params, unsafe.Pointer(cfColumn))

	var cRow int64 = int64(row)
	cfRow := C.CFNumberCreate(nil, C.kCFNumberSInt64Type, unsafe.Pointer(&cRow))
	defer C.CFRelease(C.CFTypeRef(cfRow))
	C.CFArrayAppendValue(params, unsafe.Pointer(cfRow))

	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, params, &value)
	ret := convertCFType(value)
	o, ok := ret.(*UIElement)
	if !ok {
		return nil
	}
	return o
}

func (ref *UIElement) pointForPointAttr(attribute string, p Point) Point {
	a := cfstr(attribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfPoint := C.CGPointMake(C.CGFloat(p.X), C.CGFloat(p.Y))
	axPointValue := C.AXValueCreate(C.kAXValueTypeCGPoint, unsafe.Pointer(&cfPoint))
	defer C.CFRelease(C.CFTypeRef(axPointValue))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, axPointValue, &value)
	ret := convertCFType(value)
	o, ok := ret.(Point)
	if !ok {
		return Point{0, 0}
	}
	return o
}

// LayoutPointForScreenPoint returns kAXLayoutPointForScreenPointParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) LayoutPointForScreenPoint(p Point) Point {
	return ref.pointForPointAttr(LayoutPointForScreenPointParameterizedAttribute, p)
}

func (ref *UIElement) sizeForSizeAttr(attribute string, s Size) Size {
	a := cfstr(attribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfSize := C.CGSizeMake(C.CGFloat(s.Dx), C.CGFloat(s.Dy))
	axSizeValue := C.AXValueCreate(C.kAXValueTypeCGSize, unsafe.Pointer(&cfSize))
	defer C.CFRelease(C.CFTypeRef(axSizeValue))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, axSizeValue, &value)
	ret := convertCFType(value)
	o, ok := ret.(Size)
	if !ok {
		return Size{0, 0}
	}
	return o
}

// LayoutSizeForScreenSize returns kAXLayoutSizeForScreenSizeParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) LayoutSizeForScreenSize(s Size) Size {
	return ref.sizeForSizeAttr(LayoutSizeForScreenSizeParameterizedAttribute, s)
}

// LineForIndex returns kAXLineForIndexParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) LineForIndex(index int32) int32 {
	a := cfstr(LayoutSizeForScreenSizeParameterizedAttribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfNumber := C.CFNumberCreate(nil, C.kCFNumberSInt32Type, unsafe.Pointer(&index))
	defer C.CFRelease(C.CFTypeRef(cfNumber))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, cfNumber, &value)
	ret := convertCFType(value)
	o, ok := ret.(int32)
	if !ok {
		return 0
	}
	return o
}

func (ref *UIElement) rangeForInt32Attr(attribute string, i32 int32) Range {
	a := cfstr(attribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfNumber := C.CFNumberCreate(nil, C.kCFNumberSInt32Type, unsafe.Pointer(&i32))
	defer C.CFRelease(C.CFTypeRef(cfNumber))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, cfNumber, &value)
	ret := convertCFType(value)
	o, ok := ret.(Range)
	if !ok {
		return Range{0, 0}
	}
	return o
}

// RangeForIndex returns kAXRangeForIndexParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) RangeForIndex(index int32) Range {
	return ref.rangeForInt32Attr(RangeForIndexParameterizedAttribute, index)
}

// RangeForLine returns kAXRangeForLineParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) RangeForLine(line int32) Range {
	return ref.rangeForInt32Attr(RangeForLineParameterizedAttribute, line)
}

// RangeForPosition returns kAXRangeForPositionParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) RangeForPosition(p Point) Range {
	a := cfstr(RangeForPositionParameterizedAttribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfPoint := C.CGPointMake(C.CGFloat(p.X), C.CGFloat(p.Y))
	axPointValue := C.AXValueCreate(C.kAXValueTypeCGPoint, unsafe.Pointer(&cfPoint))
	defer C.CFRelease(C.CFTypeRef(axPointValue))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, axPointValue, &value)
	ret := convertCFType(value)
	o, ok := ret.(Range)
	if !ok {
		return Range{0, 0}
	}
	return o
}

// ScreenPointForLayoutPoint returns kAXScreenPointForLayoutPointParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) ScreenPointForLayoutPoint(p Point) Point {
	return ref.pointForPointAttr(ScreenPointForLayoutPointParameterizedAttribute, p)
}

// ScreenSizeForLayoutSize returns kAXScreenSizeForLayoutSizeParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) ScreenSizeForLayoutSize(s Size) Size {
	return ref.sizeForSizeAttr(ScreenSizeForLayoutSizeParameterizedAttribute, s)
}

// StringForRange returns kAXStringForRangeParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) StringForRange(r Range) string {
	a := cfstr(StringForRangeParameterizedAttribute)
	defer C.CFRelease(C.CFTypeRef(a))
	cfRange := C.CFRangeMake(C.CFIndex(r.Location), C.CFIndex(r.Length))
	axRangeValue := C.AXValueCreate(C.kAXValueTypeCFRange, unsafe.Pointer(&cfRange))
	defer C.CFRelease(C.CFTypeRef(axRangeValue))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, axRangeValue, &value)
	ret := convertCFType(value)
	o, ok := ret.(string)
	if !ok {
		return ""
	}
	return o
}

// StyleRangeForIndex returns kAXStyleRangeForIndexParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) StyleRangeForIndex(index int32) Range {
	return ref.rangeForInt32Attr(StyleRangeForIndexParameterizedAttribute, index)
}

// AllowedValues returns kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValues() []interface{} {
	items := []interface{}{}
	ret, err := ref.attr(AllowedValuesAttribute, C.CFArrayGetTypeID())
	if err != nil {
		return items
	}
	num := int(C.CFArrayGetCount(C.CFArrayRef(ret)))
	for i := 0; i < num; i++ {
		o := C.CFTypeRef(C.CFArrayGetValueAtIndex(C.CFArrayRef(ret), C.CFIndex(i)))
		if o == nil {
			continue
		}
		C.CFRetain(o)
		t := convertCFType(o)
		if t != nil {
			items = append(items, t)
		}
	}
	return items
}

// AllowedValuesAsInt32 returns explicitly int32 casted kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValuesAsInt32() (ret []int32, err error) {
	ret = []int32{}
	err = nil
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	items := ref.AllowedValues()
	for _, item := range items {
		ret = append(ret, item.(int32))
	}
	return ret, nil
}

// AllowedValuesAsInt64 returns explicitly int64 casted kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValuesAsInt64() (ret []int64, err error) {
	ret = []int64{}
	err = nil
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	items := ref.AllowedValues()
	for _, item := range items {
		ret = append(ret, item.(int64))
	}
	return ret, nil
}

// AllowedValuesAsFloat32 returns explicitly float32 casted kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValuesAsFloat32() (ret []float32, err error) {
	ret = []float32{}
	err = nil
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	items := ref.AllowedValues()
	for _, item := range items {
		ret = append(ret, item.(float32))
	}
	return ret, nil
}

// AllowedValuesAsFloat64 returns explicitly float64 casted kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValuesAsFloat64() (ret []float64, err error) {
	ret = []float64{}
	err = nil
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	items := ref.AllowedValues()
	for _, item := range items {
		ret = append(ret, item.(float64))
	}
	return ret, nil
}

// AllowedValuesAsBool returns explicitly bool casted kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValuesAsBool() (ret []bool, err error) {
	ret = []bool{}
	err = nil
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	items := ref.AllowedValues()
	for _, item := range items {
		ret = append(ret, item.(bool))
	}
	return ret, nil
}

// AllowedValuesAsString returns explicitly string casted kAXAllowedValuesAttribute of AXUIElementRef.
func (ref *UIElement) AllowedValuesAsString() (ret []string, err error) {
	ret = []string{}
	err = nil
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	items := ref.AllowedValues()
	for _, item := range items {
		ret = append(ret, item.(string))
	}
	return ret, nil
}

// RTFForRange returns kAXRTFForRangeParameterizedAttribute of AXUIElementRef.
func (ref *UIElement) RTFForRange(r Range) []byte {
	a := cfstr(RTFForRangeParameterizedAttribute)
	defer C.CFRelease(C.CFTypeRef(a))

	cfRange := C.CFRangeMake(C.CFIndex(r.Location), C.CFIndex(r.Length))
	axRangeValue := C.AXValueCreate(C.kAXValueTypeCFRange, unsafe.Pointer(&cfRange))
	defer C.CFRelease(C.CFTypeRef(axRangeValue))
	var value C.CFTypeRef
	C.AXUIElementCopyParameterizedAttributeValue(ref.obj, a, axRangeValue, &value)
	if value == nil {
		return []byte{}
	}
	defer C.CFRelease(value)
	if C.CFGetTypeID(value) != C.CFDataGetTypeID() {
		return []byte{}
	}
	length := int(C.CFDataGetLength(C.CFDataRef(value)))
	data := make([]byte, length)
	ptr := uintptr(unsafe.Pointer(C.CFDataGetBytePtr(C.CFDataRef(value))))
	for i := 0; i < length; i++ {
		data[i] = byte(*(*C.UInt8)(unsafe.Pointer(ptr)))
		ptr = ptr + 1
	}
	return data
}

//TODO:
// AttributedStringForRangeParameterizedAttribute  = "AXAttributedStringForRange"   // CFAttributedString

// Automatically generated funcs.

// AMPMField returns kAXAMPMFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityampmfieldattribute
func (ref *UIElement) AMPMField() *UIElement {
	ret, _ := ref.UIElementAttr(AMPMFieldAttribute)
	return ret
}

// AlternateUIVisible returns kAXAlternateUIVisibleAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityalternateuivisibleattribute
func (ref *UIElement) IsAlternateUIVisible() bool {
	ret, _ := ref.BoolAttr(AlternateUIVisibleAttribute)
	return ret
}

// CancelButton returns kAXCancelButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycancelbuttonattribute
func (ref *UIElement) CancelButton() *UIElement {
	ret, _ := ref.UIElementAttr(CancelButtonAttribute)
	return ret
}

// Children returns kAXChildrenAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitychildrenattribute
func (ref *UIElement) Children() []*UIElement {
	return ref.SliceOfUIElementAttr(ChildrenAttribute)
}

// ClearButton returns kAXClearButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityclearbuttonattribute
func (ref *UIElement) ClearButton() *UIElement {
	ret, _ := ref.UIElementAttr(ClearButtonAttribute)
	return ret
}

// CloseButton returns kAXCloseButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityclosebuttonattribute
func (ref *UIElement) CloseButton() *UIElement {
	ret, _ := ref.UIElementAttr(CloseButtonAttribute)
	return ret
}

// ColumnCount returns kAXColumnCountAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycolumncountattribute
func (ref *UIElement) ColumnCount() int64 {
	ret, _ := ref.Int64Attr(ColumnCountAttribute)
	return ret
}

// ColumnHeaderUIElements returns kAXColumnHeaderUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycolumnheaderuielementsattribute
func (ref *UIElement) ColumnHeaderUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(ColumnHeaderUIElementsAttribute)
}

// ColumnIndexRange returns kAXColumnIndexRangeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycolumnindexrangeattribute
func (ref *UIElement) ColumnIndexRange() Range {
	ret, _ := ref.RangeAttr(ColumnIndexRangeAttribute)
	return ret
}

// ColumnTitle returns kAXColumnTitleAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycolumntitleattribute
func (ref *UIElement) ColumnTitle() []string {
	return ref.SliceOfStringAttr(ColumnTitleAttribute)
}

// ColumnTitles returns kAXColumnTitlesAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycolumntitlesattribute
func (ref *UIElement) ColumnTitles() []string {
	return ref.SliceOfStringAttr(ColumnTitlesAttribute)
}

// Columns returns kAXColumnsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycolumnsattribute
func (ref *UIElement) Columns() []*UIElement {
	return ref.SliceOfUIElementAttr(ColumnsAttribute)
}

// Contents returns kAXContentsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycontentsattribute
func (ref *UIElement) Contents() []*UIElement {
	return ref.SliceOfUIElementAttr(ContentsAttribute)
}

// CriticalValue returns kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValue() interface{} {
	return ref.attrAny(CriticalValueAttribute)
}

// CriticalValueAsInt32 returns explicitly int32 casted kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValueAsInt32() (int32, error) {
	return ref.Int32Attr(CriticalValueAttribute)
}

// CriticalValueAsInt64 returns explicitly int64 casted kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValueAsInt64() (int64, error) {
	return ref.Int64Attr(CriticalValueAttribute)
}

// CriticalValueAsFloat32 returns explicitly float32 casted kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValueAsFloat32() (float32, error) {
	return ref.Float32Attr(CriticalValueAttribute)
}

// CriticalValueAsFloat64 returns explicitly float64 casted kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValueAsFloat64() (float64, error) {
	return ref.Float64Attr(CriticalValueAttribute)
}

// CriticalValueAsBool returns explicitly bool casted kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValueAsBool() (bool, error) {
	return ref.BoolAttr(CriticalValueAttribute)
}

// CriticalValueAsString returns explicitly string casted kAXCriticalValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitycriticalvalueattribute
func (ref *UIElement) CriticalValueAsString() (string, error) {
	return ref.StringAttr(CriticalValueAttribute)
}

// DayField returns kAXDayFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydayfieldattribute
func (ref *UIElement) DayField() *UIElement {
	ret, _ := ref.UIElementAttr(DayFieldAttribute)
	return ret
}

// DecrementButton returns kAXDecrementButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydecrementbuttonattribute
func (ref *UIElement) DecrementButton() *UIElement {
	ret, _ := ref.UIElementAttr(DecrementButtonAttribute)
	return ret
}

// DefaultButton returns kAXDefaultButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydefaultbuttonattribute
func (ref *UIElement) DefaultButton() *UIElement {
	ret, _ := ref.UIElementAttr(DefaultButtonAttribute)
	return ret
}

// Description returns kAXDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydescriptionattribute
func (ref *UIElement) Description() string {
	ret, _ := ref.StringAttr(DescriptionAttribute)
	return ret
}

// DisclosedByRow returns kAXDisclosedByRowAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydisclosedbyrowattribute
func (ref *UIElement) DisclosedByRow() *UIElement {
	ret, _ := ref.UIElementAttr(DisclosedByRowAttribute)
	return ret
}

// DisclosedRows returns kAXDisclosedRowsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydisclosedrowsattribute
func (ref *UIElement) DisclosedRows() []*UIElement {
	return ref.SliceOfUIElementAttr(DisclosedRowsAttribute)
}

// Disclosing returns kAXDisclosingAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydisclosingattribute
func (ref *UIElement) IsDisclosing() bool {
	ret, _ := ref.BoolAttr(DisclosingAttribute)
	return ret
}

// DisclosureLevel returns kAXDisclosureLevelAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydisclosurelevelattribute
func (ref *UIElement) DisclosureLevel() int64 {
	ret, _ := ref.Int64Attr(DisclosureLevelAttribute)
	return ret
}

// Document returns kAXDocumentAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitydocumentattribute
func (ref *UIElement) Document() string {
	ret, _ := ref.StringAttr(DocumentAttribute)
	return ret
}

// Edited returns kAXEditedAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityeditedattribute
func (ref *UIElement) IsEdited() bool {
	ret, _ := ref.BoolAttr(EditedAttribute)
	return ret
}

// ElementBusy returns kAXElementBusyAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityelementbusyattribute
func (ref *UIElement) IsElementBusy() bool {
	ret, _ := ref.BoolAttr(ElementBusyAttribute)
	return ret
}

// Enabled returns kAXEnabledAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityenabledattribute
func (ref *UIElement) IsEnabled() bool {
	ret, _ := ref.BoolAttr(EnabledAttribute)
	return ret
}

// Expanded returns kAXExpandedAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityexpandedattribute
func (ref *UIElement) IsExpanded() bool {
	ret, _ := ref.BoolAttr(ExpandedAttribute)
	return ret
}

// ExtrasMenuBar returns kAXExtrasMenuBarAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityextrasmenubarattribute
func (ref *UIElement) ExtrasMenuBar() *UIElement {
	ret, _ := ref.UIElementAttr(ExtrasMenuBarAttribute)
	return ret
}

// Filename returns kAXFilenameAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfilenameattribute
func (ref *UIElement) Filename() string {
	ret, _ := ref.StringAttr(FilenameAttribute)
	return ret
}

// FocusedApplication returns kAXFocusedApplicationAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfocusedapplicationattribute
func (ref *UIElement) FocusedApplication() *UIElement {
	ret, _ := ref.UIElementAttr(FocusedApplicationAttribute)
	return ret
}

// Focused returns kAXFocusedAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfocusedattribute
func (ref *UIElement) IsFocused() bool {
	ret, _ := ref.BoolAttr(FocusedAttribute)
	return ret
}

// FocusedUIElement returns kAXFocusedUIElementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfocuseduielementattribute
func (ref *UIElement) FocusedUIElement() *UIElement {
	ret, _ := ref.UIElementAttr(FocusedUIElementAttribute)
	return ret
}

// FocusedWindow returns kAXFocusedWindowAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfocusedwindowattribute
func (ref *UIElement) FocusedWindow() *UIElement {
	ret, _ := ref.UIElementAttr(FocusedWindowAttribute)
	return ret
}

// Frontmost returns kAXFrontmostAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfrontmostattribute
func (ref *UIElement) IsFrontmost() bool {
	ret, _ := ref.BoolAttr(FrontmostAttribute)
	return ret
}

// FullScreenButton returns kAXFullScreenButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityfullscreenbuttonattribute
func (ref *UIElement) FullScreenButton() *UIElement {
	ret, _ := ref.UIElementAttr(FullScreenButtonAttribute)
	return ret
}

// GrowArea returns kAXGrowAreaAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitygrowareaattribute
func (ref *UIElement) GrowArea() *UIElement {
	ret, _ := ref.UIElementAttr(GrowAreaAttribute)
	return ret
}

// Handles returns kAXHandlesAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhandlesattribute
func (ref *UIElement) Handles() []*UIElement {
	return ref.SliceOfUIElementAttr(HandlesAttribute)
}

// Header returns kAXHeaderAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityheaderattribute
func (ref *UIElement) Header() *UIElement {
	ret, _ := ref.UIElementAttr(HeaderAttribute)
	return ret
}

// Help returns kAXHelpAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhelpattribute
func (ref *UIElement) Help() string {
	ret, _ := ref.StringAttr(HelpAttribute)
	return ret
}

// Hidden returns kAXHiddenAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhiddenattribute
func (ref *UIElement) IsHidden() bool {
	ret, _ := ref.BoolAttr(HiddenAttribute)
	return ret
}

// HorizontalScrollBar returns kAXHorizontalScrollBarAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhorizontalscrollbarattribute
func (ref *UIElement) HorizontalScrollBar() *UIElement {
	ret, _ := ref.UIElementAttr(HorizontalScrollBarAttribute)
	return ret
}

// HorizontalUnitDescription returns kAXHorizontalUnitDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhorizontalunitdescriptionattribute
func (ref *UIElement) HorizontalUnitDescription() string {
	ret, _ := ref.StringAttr(HorizontalUnitDescriptionAttribute)
	return ret
}

// HorizontalUnits returns kAXHorizontalUnitsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhorizontalunitsattribute
func (ref *UIElement) HorizontalUnits() string {
	ret, _ := ref.StringAttr(HorizontalUnitsAttribute)
	return ret
}

// HourField returns kAXHourFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityhourfieldattribute
func (ref *UIElement) HourField() *UIElement {
	ret, _ := ref.UIElementAttr(HourFieldAttribute)
	return ret
}

// Identifier returns kAXIdentifierAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityidentifierattribute
func (ref *UIElement) Identifier() string {
	ret, _ := ref.StringAttr(IdentifierAttribute)
	return ret
}

// IncrementButton returns kAXIncrementButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityincrementbuttonattribute
func (ref *UIElement) IncrementButton() *UIElement {
	ret, _ := ref.UIElementAttr(IncrementButtonAttribute)
	return ret
}

// Incrementor returns kAXIncrementorAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityincrementorattribute
func (ref *UIElement) Incrementor() *UIElement {
	ret, _ := ref.UIElementAttr(IncrementorAttribute)
	return ret
}

// Index returns kAXIndexAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityindexattribute
func (ref *UIElement) Index() int64 {
	ret, _ := ref.Int64Attr(IndexAttribute)
	return ret
}

// InsertionPointLineNumber returns kAXInsertionPointLineNumberAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityinsertionpointlinenumberattribute
func (ref *UIElement) InsertionPointLineNumber() int64 {
	ret, _ := ref.Int64Attr(InsertionPointLineNumberAttribute)
	return ret
}

// IsApplicationRunning returns kAXIsApplicationRunningAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityisapplicationrunningattribute
func (ref *UIElement) IsApplicationRunning() bool {
	ret, _ := ref.BoolAttr(IsApplicationRunningAttribute)
	return ret
}

// IsEditable returns kAXIsEditableAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityiseditableattribute
func (ref *UIElement) IsEditable() bool {
	ret, _ := ref.BoolAttr(IsEditableAttribute)
	return ret
}

// LabelUIElements returns kAXLabelUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitylabeluielementsattribute
func (ref *UIElement) LabelUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(LabelUIElementsAttribute)
}

// LabelValue returns kAXLabelValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitylabelvalueattribute
func (ref *UIElement) LabelValue() float32 {
	ret, _ := ref.Float32Attr(LabelValueAttribute)
	return ret
}

// LinkedUIElements returns kAXLinkedUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitylinkeduielementsattribute
func (ref *UIElement) LinkedUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(LinkedUIElementsAttribute)
}

// Main returns kAXMainAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymainattribute
func (ref *UIElement) IsMain() bool {
	ret, _ := ref.BoolAttr(MainAttribute)
	return ret
}

// MainWindow returns kAXMainWindowAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymainwindowattribute
func (ref *UIElement) MainWindow() *UIElement {
	ret, _ := ref.UIElementAttr(MainWindowAttribute)
	return ret
}

// MarkerType returns kAXMarkerTypeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymarkertypeattribute
func (ref *UIElement) MarkerType() string {
	ret, _ := ref.StringAttr(MarkerTypeAttribute)
	return ret
}

// MarkerTypeDescription returns kAXMarkerTypeDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymarkertypedescriptionattribute
func (ref *UIElement) MarkerTypeDescription() string {
	ret, _ := ref.StringAttr(MarkerTypeDescriptionAttribute)
	return ret
}

// MarkerUIElements returns kAXMarkerUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymarkeruielementsattribute
func (ref *UIElement) MarkerUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(MarkerUIElementsAttribute)
}

// MatteContentUIElement returns kAXMatteContentUIElementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymattecontentuielementattribute
func (ref *UIElement) MatteContentUIElement() *UIElement {
	ret, _ := ref.UIElementAttr(MatteContentUIElementAttribute)
	return ret
}

// MatteHole returns kAXMatteHoleAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymatteholeattribute
func (ref *UIElement) MatteHole() Rect {
	ret, _ := ref.RectAttr(MatteHoleAttribute)
	return ret
}

// MaxValue returns kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValue() interface{} {
	return ref.attrAny(MaxValueAttribute)
}

// MaxValueAsInt32 returns explicitly int32 casted kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValueAsInt32() (int32, error) {
	return ref.Int32Attr(MaxValueAttribute)
}

// MaxValueAsInt64 returns explicitly int64 casted kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValueAsInt64() (int64, error) {
	return ref.Int64Attr(MaxValueAttribute)
}

// MaxValueAsFloat32 returns explicitly float32 casted kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValueAsFloat32() (float32, error) {
	return ref.Float32Attr(MaxValueAttribute)
}

// MaxValueAsFloat64 returns explicitly float64 casted kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValueAsFloat64() (float64, error) {
	return ref.Float64Attr(MaxValueAttribute)
}

// MaxValueAsBool returns explicitly bool casted kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValueAsBool() (bool, error) {
	return ref.BoolAttr(MaxValueAttribute)
}

// MaxValueAsString returns explicitly string casted kAXMaxValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymaxvalueattribute
func (ref *UIElement) MaxValueAsString() (string, error) {
	return ref.StringAttr(MaxValueAttribute)
}

// MenuBar returns kAXMenuBarAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenubarattribute
func (ref *UIElement) MenuBar() *UIElement {
	ret, _ := ref.UIElementAttr(MenuBarAttribute)
	return ret
}

// MenuItemCmdChar returns kAXMenuItemCmdCharAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenuitemcmdcharattribute
func (ref *UIElement) MenuItemCmdChar() string {
	ret, _ := ref.StringAttr(MenuItemCmdCharAttribute)
	return ret
}

// MenuItemCmdGlyph returns kAXMenuItemCmdGlyphAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenuitemcmdglyphattribute
func (ref *UIElement) MenuItemCmdGlyph() int64 {
	ret, _ := ref.Int64Attr(MenuItemCmdGlyphAttribute)
	return ret
}

// MenuItemCmdModifiers returns kAXMenuItemCmdModifiersAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenuitemcmdmodifiersattribute
func (ref *UIElement) MenuItemCmdModifiers() int64 {
	ret, _ := ref.Int64Attr(MenuItemCmdModifiersAttribute)
	return ret
}

// MenuItemCmdVirtualKey returns kAXMenuItemCmdVirtualKeyAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenuitemcmdvirtualkeyattribute
func (ref *UIElement) MenuItemCmdVirtualKey() int64 {
	ret, _ := ref.Int64Attr(MenuItemCmdVirtualKeyAttribute)
	return ret
}

// MenuItemMarkChar returns kAXMenuItemMarkCharAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenuitemmarkcharattribute
func (ref *UIElement) MenuItemMarkChar() string {
	ret, _ := ref.StringAttr(MenuItemMarkCharAttribute)
	return ret
}

// MenuItemPrimaryUIElement returns kAXMenuItemPrimaryUIElementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymenuitemprimaryuielementattribute
func (ref *UIElement) MenuItemPrimaryUIElement() *UIElement {
	ret, _ := ref.UIElementAttr(MenuItemPrimaryUIElementAttribute)
	return ret
}

// MinValue returns kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValue() interface{} {
	return ref.attrAny(MinValueAttribute)
}

// MinValueAsInt32 returns explicitly int32 casted kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValueAsInt32() (int32, error) {
	return ref.Int32Attr(MinValueAttribute)
}

// MinValueAsInt64 returns explicitly int64 casted kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValueAsInt64() (int64, error) {
	return ref.Int64Attr(MinValueAttribute)
}

// MinValueAsFloat32 returns explicitly float32 casted kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValueAsFloat32() (float32, error) {
	return ref.Float32Attr(MinValueAttribute)
}

// MinValueAsFloat64 returns explicitly float64 casted kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValueAsFloat64() (float64, error) {
	return ref.Float64Attr(MinValueAttribute)
}

// MinValueAsBool returns explicitly bool casted kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValueAsBool() (bool, error) {
	return ref.BoolAttr(MinValueAttribute)
}

// MinValueAsString returns explicitly string casted kAXMinValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminvalueattribute
func (ref *UIElement) MinValueAsString() (string, error) {
	return ref.StringAttr(MinValueAttribute)
}

// MinimizeButton returns kAXMinimizeButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminimizebuttonattribute
func (ref *UIElement) MinimizeButton() *UIElement {
	ret, _ := ref.UIElementAttr(MinimizeButtonAttribute)
	return ret
}

// Minimized returns kAXMinimizedAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminimizedattribute
func (ref *UIElement) IsMinimized() bool {
	ret, _ := ref.BoolAttr(MinimizedAttribute)
	return ret
}

// MinuteField returns kAXMinuteFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityminutefieldattribute
func (ref *UIElement) MinuteField() *UIElement {
	ret, _ := ref.UIElementAttr(MinuteFieldAttribute)
	return ret
}

// Modal returns kAXModalAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymodalattribute
func (ref *UIElement) IsModal() bool {
	ret, _ := ref.BoolAttr(ModalAttribute)
	return ret
}

// MonthField returns kAXMonthFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitymonthfieldattribute
func (ref *UIElement) MonthField() *UIElement {
	ret, _ := ref.UIElementAttr(MonthFieldAttribute)
	return ret
}

// NextContents returns kAXNextContentsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitynextcontentsattribute
func (ref *UIElement) NextContents() []*UIElement {
	return ref.SliceOfUIElementAttr(NextContentsAttribute)
}

// NumberOfCharacters returns kAXNumberOfCharactersAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitynumberofcharactersattribute
func (ref *UIElement) NumberOfCharacters() int64 {
	ret, _ := ref.Int64Attr(NumberOfCharactersAttribute)
	return ret
}

// OrderedByRow returns kAXOrderedByRowAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityorderedbyrowattribute
func (ref *UIElement) IsOrderedByRow() bool {
	ret, _ := ref.BoolAttr(OrderedByRowAttribute)
	return ret
}

// Orientation returns kAXOrientationAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityorientationattribute
func (ref *UIElement) Orientation() string {
	ret, _ := ref.StringAttr(OrientationAttribute)
	return ret
}

// OverflowButton returns kAXOverflowButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityoverflowbuttonattribute
func (ref *UIElement) OverflowButton() *UIElement {
	ret, _ := ref.UIElementAttr(OverflowButtonAttribute)
	return ret
}

// Parent returns kAXParentAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityparentattribute
func (ref *UIElement) Parent() *UIElement {
	ret, _ := ref.UIElementAttr(ParentAttribute)
	return ret
}

// PlaceholderValue returns kAXPlaceholderValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityplaceholdervalueattribute
func (ref *UIElement) PlaceholderValue() string {
	ret, _ := ref.StringAttr(PlaceholderValueAttribute)
	return ret
}

// Position returns kAXPositionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitypositionattribute
func (ref *UIElement) Position() Point {
	ret, _ := ref.PointAttr(PositionAttribute)
	return ret
}

// PreviousContents returns kAXPreviousContentsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitypreviouscontentsattribute
func (ref *UIElement) PreviousContents() []*UIElement {
	return ref.SliceOfUIElementAttr(PreviousContentsAttribute)
}

// Proxy returns kAXProxyAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityproxyattribute
func (ref *UIElement) Proxy() *UIElement {
	ret, _ := ref.UIElementAttr(ProxyAttribute)
	return ret
}

// Role returns kAXRoleAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityroleattribute
func (ref *UIElement) Role() string {
	ret, _ := ref.StringAttr(RoleAttribute)
	return ret
}

// RoleDescription returns kAXRoleDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityroledescriptionattribute
func (ref *UIElement) RoleDescription() string {
	ret, _ := ref.StringAttr(RoleDescriptionAttribute)
	return ret
}

// RowCount returns kAXRowCountAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityrowcountattribute
func (ref *UIElement) RowCount() int64 {
	ret, _ := ref.Int64Attr(RowCountAttribute)
	return ret
}

// RowHeaderUIElements returns kAXRowHeaderUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityrowheaderuielementsattribute
func (ref *UIElement) RowHeaderUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(RowHeaderUIElementsAttribute)
}

// RowIndexRange returns kAXRowIndexRangeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityrowindexrangeattribute
func (ref *UIElement) RowIndexRange() Range {
	ret, _ := ref.RangeAttr(RowIndexRangeAttribute)
	return ret
}

// Rows returns kAXRowsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityrowsattribute
func (ref *UIElement) Rows() []*UIElement {
	return ref.SliceOfUIElementAttr(RowsAttribute)
}

// SearchButton returns kAXSearchButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysearchbuttonattribute
func (ref *UIElement) SearchButton() *UIElement {
	ret, _ := ref.UIElementAttr(SearchButtonAttribute)
	return ret
}

// SecondField returns kAXSecondFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysecondfieldattribute
func (ref *UIElement) SecondField() *UIElement {
	ret, _ := ref.UIElementAttr(SecondFieldAttribute)
	return ret
}

// Selected returns kAXSelectedAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedattribute
func (ref *UIElement) IsSelected() bool {
	ret, _ := ref.BoolAttr(SelectedAttribute)
	return ret
}

// SelectedCells returns kAXSelectedCellsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedcellsattribute
func (ref *UIElement) SelectedCells() []*UIElement {
	return ref.SliceOfUIElementAttr(SelectedCellsAttribute)
}

// SelectedChildren returns kAXSelectedChildrenAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedchildrenattribute
func (ref *UIElement) SelectedChildren() []*UIElement {
	return ref.SliceOfUIElementAttr(SelectedChildrenAttribute)
}

// SelectedColumns returns kAXSelectedColumnsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedcolumnsattribute
func (ref *UIElement) SelectedColumns() []*UIElement {
	return ref.SliceOfUIElementAttr(SelectedColumnsAttribute)
}

// SelectedRows returns kAXSelectedRowsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedrowsattribute
func (ref *UIElement) SelectedRows() []*UIElement {
	return ref.SliceOfUIElementAttr(SelectedRowsAttribute)
}

// SelectedText returns kAXSelectedTextAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedtextattribute
func (ref *UIElement) SelectedText() string {
	ret, _ := ref.StringAttr(SelectedTextAttribute)
	return ret
}

// SelectedTextRange returns kAXSelectedTextRangeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedtextrangeattribute
func (ref *UIElement) SelectedTextRange() Range {
	ret, _ := ref.RangeAttr(SelectedTextRangeAttribute)
	return ret
}

// SelectedTextRanges returns kAXSelectedTextRangesAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityselectedtextrangesattribute
func (ref *UIElement) SelectedTextRanges() []Range {
	return ref.SliceOfRangeAttr(SelectedTextRangesAttribute)
}

// ServesAsTitleForUIElements returns kAXServesAsTitleForUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityservesastitleforuielementsattribute
func (ref *UIElement) ServesAsTitleForUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(ServesAsTitleForUIElementsAttribute)
}

// SharedCharacterRange returns kAXSharedCharacterRangeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysharedcharacterrangeattribute
func (ref *UIElement) SharedCharacterRange() Range {
	ret, _ := ref.RangeAttr(SharedCharacterRangeAttribute)
	return ret
}

// SharedFocusElements returns kAXSharedFocusElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysharedfocuselementsattribute
func (ref *UIElement) SharedFocusElements() []*UIElement {
	return ref.SliceOfUIElementAttr(SharedFocusElementsAttribute)
}

// SharedTextUIElements returns kAXSharedTextUIElementsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysharedtextuielementsattribute
func (ref *UIElement) SharedTextUIElements() []*UIElement {
	return ref.SliceOfUIElementAttr(SharedTextUIElementsAttribute)
}

// ShownMenuUIElement returns kAXShownMenuUIElementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityshownmenuuielementattribute
func (ref *UIElement) ShownMenuUIElement() *UIElement {
	ret, _ := ref.UIElementAttr(ShownMenuUIElementAttribute)
	return ret
}

// Size returns kAXSizeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysizeattribute
func (ref *UIElement) Size() Size {
	ret, _ := ref.SizeAttr(SizeAttribute)
	return ret
}

// SortDirection returns kAXSortDirectionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysortdirectionattribute
func (ref *UIElement) SortDirection() string {
	ret, _ := ref.StringAttr(SortDirectionAttribute)
	return ret
}

// Splitters returns kAXSplittersAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysplittersattribute
func (ref *UIElement) Splitters() []*UIElement {
	return ref.SliceOfUIElementAttr(SplittersAttribute)
}

// Subrole returns kAXSubroleAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitysubroleattribute
func (ref *UIElement) Subrole() string {
	ret, _ := ref.StringAttr(SubroleAttribute)
	return ret
}

// Tabs returns kAXTabsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitytabsattribute
func (ref *UIElement) Tabs() []*UIElement {
	return ref.SliceOfUIElementAttr(TabsAttribute)
}

// Text returns kAXTextAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitytextattribute
func (ref *UIElement) Text() string {
	ret, _ := ref.StringAttr(TextAttribute)
	return ret
}

// Title returns kAXTitleAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitytitleattribute
func (ref *UIElement) Title() string {
	ret, _ := ref.StringAttr(TitleAttribute)
	return ret
}

// TitleUIElement returns kAXTitleUIElementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitytitleuielementattribute
func (ref *UIElement) TitleUIElement() *UIElement {
	ret, _ := ref.UIElementAttr(TitleUIElementAttribute)
	return ret
}

// ToolbarButton returns kAXToolbarButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitytoolbarbuttonattribute
func (ref *UIElement) ToolbarButton() *UIElement {
	ret, _ := ref.UIElementAttr(ToolbarButtonAttribute)
	return ret
}

// TopLevelUIElement returns kAXTopLevelUIElementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitytopleveluielementattribute
func (ref *UIElement) TopLevelUIElement() *UIElement {
	ret, _ := ref.UIElementAttr(TopLevelUIElementAttribute)
	return ret
}

// URL returns kAXURLAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityurlattribute
func (ref *UIElement) URL() *url.URL {
	ret, _ := ref.URLAttr(URLAttribute)
	return ret
}

// UnitDescription returns kAXUnitDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityunitdescriptionattribute
func (ref *UIElement) UnitDescription() string {
	ret, _ := ref.StringAttr(UnitDescriptionAttribute)
	return ret
}

// Units returns kAXUnitsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityunitsattribute
func (ref *UIElement) Units() string {
	ret, _ := ref.StringAttr(UnitsAttribute)
	return ret
}

// Value returns kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) Value() interface{} {
	return ref.attrAny(ValueAttribute)
}

// ValueAsInt32 returns explicitly int32 casted kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) ValueAsInt32() (int32, error) {
	return ref.Int32Attr(ValueAttribute)
}

// ValueAsInt64 returns explicitly int64 casted kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) ValueAsInt64() (int64, error) {
	return ref.Int64Attr(ValueAttribute)
}

// ValueAsFloat32 returns explicitly float32 casted kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) ValueAsFloat32() (float32, error) {
	return ref.Float32Attr(ValueAttribute)
}

// ValueAsFloat64 returns explicitly float64 casted kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) ValueAsFloat64() (float64, error) {
	return ref.Float64Attr(ValueAttribute)
}

// ValueAsBool returns explicitly bool casted kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) ValueAsBool() (bool, error) {
	return ref.BoolAttr(ValueAttribute)
}

// ValueAsString returns explicitly string casted kAXValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueattribute
func (ref *UIElement) ValueAsString() (string, error) {
	return ref.StringAttr(ValueAttribute)
}

// ValueDescription returns kAXValueDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvaluedescriptionattribute
func (ref *UIElement) ValueDescription() string {
	ret, _ := ref.StringAttr(ValueDescriptionAttribute)
	return ret
}

// ValueIncrement returns kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrement() interface{} {
	return ref.attrAny(ValueIncrementAttribute)
}

// ValueIncrementAsInt32 returns explicitly int32 casted kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrementAsInt32() (int32, error) {
	return ref.Int32Attr(ValueIncrementAttribute)
}

// ValueIncrementAsInt64 returns explicitly int64 casted kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrementAsInt64() (int64, error) {
	return ref.Int64Attr(ValueIncrementAttribute)
}

// ValueIncrementAsFloat32 returns explicitly float32 casted kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrementAsFloat32() (float32, error) {
	return ref.Float32Attr(ValueIncrementAttribute)
}

// ValueIncrementAsFloat64 returns explicitly float64 casted kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrementAsFloat64() (float64, error) {
	return ref.Float64Attr(ValueIncrementAttribute)
}

// ValueIncrementAsBool returns explicitly bool casted kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrementAsBool() (bool, error) {
	return ref.BoolAttr(ValueIncrementAttribute)
}

// ValueIncrementAsString returns explicitly string casted kAXValueIncrementAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvalueincrementattribute
func (ref *UIElement) ValueIncrementAsString() (string, error) {
	return ref.StringAttr(ValueIncrementAttribute)
}

// VerticalScrollBar returns kAXVerticalScrollBarAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityverticalscrollbarattribute
func (ref *UIElement) VerticalScrollBar() *UIElement {
	ret, _ := ref.UIElementAttr(VerticalScrollBarAttribute)
	return ret
}

// VerticalUnitDescription returns kAXVerticalUnitDescriptionAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityverticalunitdescriptionattribute
func (ref *UIElement) VerticalUnitDescription() string {
	ret, _ := ref.StringAttr(VerticalUnitDescriptionAttribute)
	return ret
}

// VerticalUnits returns kAXVerticalUnitsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityverticalunitsattribute
func (ref *UIElement) VerticalUnits() string {
	ret, _ := ref.StringAttr(VerticalUnitsAttribute)
	return ret
}

// VisibleCells returns kAXVisibleCellsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvisiblecellsattribute
func (ref *UIElement) VisibleCells() []*UIElement {
	return ref.SliceOfUIElementAttr(VisibleCellsAttribute)
}

// VisibleCharacterRange returns kAXVisibleCharacterRangeAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvisiblecharacterrangeattribute
func (ref *UIElement) VisibleCharacterRange() Range {
	ret, _ := ref.RangeAttr(VisibleCharacterRangeAttribute)
	return ret
}

// VisibleChildren returns kAXVisibleChildrenAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvisiblechildrenattribute
func (ref *UIElement) VisibleChildren() []*UIElement {
	return ref.SliceOfUIElementAttr(VisibleChildrenAttribute)
}

// VisibleColumns returns kAXVisibleColumnsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvisiblecolumnsattribute
func (ref *UIElement) VisibleColumns() []*UIElement {
	return ref.SliceOfUIElementAttr(VisibleColumnsAttribute)
}

// VisibleRows returns kAXVisibleRowsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvisiblerowsattribute
func (ref *UIElement) VisibleRows() []*UIElement {
	return ref.SliceOfUIElementAttr(VisibleRowsAttribute)
}

// VisibleText returns kAXVisibleTextAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityvisibletextattribute
func (ref *UIElement) VisibleText() string {
	ret, _ := ref.StringAttr(VisibleTextAttribute)
	return ret
}

// WarningValue returns kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValue() interface{} {
	return ref.attrAny(WarningValueAttribute)
}

// WarningValueAsInt32 returns explicitly int32 casted kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValueAsInt32() (int32, error) {
	return ref.Int32Attr(WarningValueAttribute)
}

// WarningValueAsInt64 returns explicitly int64 casted kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValueAsInt64() (int64, error) {
	return ref.Int64Attr(WarningValueAttribute)
}

// WarningValueAsFloat32 returns explicitly float32 casted kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValueAsFloat32() (float32, error) {
	return ref.Float32Attr(WarningValueAttribute)
}

// WarningValueAsFloat64 returns explicitly float64 casted kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValueAsFloat64() (float64, error) {
	return ref.Float64Attr(WarningValueAttribute)
}

// WarningValueAsBool returns explicitly bool casted kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValueAsBool() (bool, error) {
	return ref.BoolAttr(WarningValueAttribute)
}

// WarningValueAsString returns explicitly string casted kAXWarningValueAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywarningvalueattribute
func (ref *UIElement) WarningValueAsString() (string, error) {
	return ref.StringAttr(WarningValueAttribute)
}

// Window returns kAXWindowAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywindowattribute
func (ref *UIElement) Window() *UIElement {
	ret, _ := ref.UIElementAttr(WindowAttribute)
	return ret
}

// Windows returns kAXWindowsAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilitywindowsattribute
func (ref *UIElement) Windows() []*UIElement {
	return ref.SliceOfUIElementAttr(WindowsAttribute)
}

// YearField returns kAXYearFieldAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityyearfieldattribute
func (ref *UIElement) YearField() *UIElement {
	ret, _ := ref.UIElementAttr(YearFieldAttribute)
	return ret
}

// ZoomButton returns kAXZoomButtonAttribute of AXUIElementRef. See also https://developer.apple.com/reference/appkit/nsaccessibilityzoombuttonattribute
func (ref *UIElement) ZoomButton() *UIElement {
	ret, _ := ref.UIElementAttr(ZoomButtonAttribute)
	return ret
}
