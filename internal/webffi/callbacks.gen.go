/*
------------------------------------------------------------------------------
//   This code was generated by template ffi_wrapper.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_wrapper.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package webffi

import (
	"syscall/js"
)

func gdspxOnEngineStart(this js.Value, args []js.Value) interface{} {
	if callbacks.OnEngineStart == nil {
		return nil
	}
	callbacks.OnEngineStart()
	return nil
}
func gdspxOnEngineUpdate(this js.Value, args []js.Value) interface{} {
	if callbacks.OnEngineUpdate == nil {
		return nil
	}
	arg0 := JsToGdFloat(args[0])
	callbacks.OnEngineUpdate(arg0)
	return nil
}
func gdspxOnEngineFixedUpdate(this js.Value, args []js.Value) interface{} {
	if callbacks.OnEngineFixedUpdate == nil {
		return nil
	}
	arg0 := JsToGdFloat(args[0])
	callbacks.OnEngineFixedUpdate(arg0)
	return nil
}
func gdspxOnEngineDestroy(this js.Value, args []js.Value) interface{} {
	if callbacks.OnEngineDestroy == nil {
		return nil
	}
	callbacks.OnEngineDestroy()
	return nil
}
func gdspxOnSceneSpriteInstantiated(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSceneSpriteInstantiated == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	arg1 := JsToGdString(args[1])
	callbacks.OnSceneSpriteInstantiated(arg0, arg1)
	return nil
}
func gdspxOnSpriteReady(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteReady == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteReady(arg0)
	return nil
}
func gdspxOnSpriteUpdated(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteUpdated == nil {
		return nil
	}
	arg0 := JsToGdFloat(args[0])
	callbacks.OnSpriteUpdated(arg0)
	return nil
}
func gdspxOnSpriteFixedUpdated(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteFixedUpdated == nil {
		return nil
	}
	arg0 := JsToGdFloat(args[0])
	callbacks.OnSpriteFixedUpdated(arg0)
	return nil
}
func gdspxOnSpriteDestroyed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteDestroyed == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteDestroyed(arg0)
	return nil
}
func gdspxOnSpriteFramesSetChanged(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteFramesSetChanged == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteFramesSetChanged(arg0)
	return nil
}
func gdspxOnSpriteAnimationChanged(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteAnimationChanged == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteAnimationChanged(arg0)
	return nil
}
func gdspxOnSpriteFrameChanged(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteFrameChanged == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteFrameChanged(arg0)
	return nil
}
func gdspxOnSpriteAnimationLooped(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteAnimationLooped == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteAnimationLooped(arg0)
	return nil
}
func gdspxOnSpriteAnimationFinished(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteAnimationFinished == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteAnimationFinished(arg0)
	return nil
}
func gdspxOnSpriteVfxFinished(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteVfxFinished == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteVfxFinished(arg0)
	return nil
}
func gdspxOnSpriteScreenExited(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteScreenExited == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteScreenExited(arg0)
	return nil
}
func gdspxOnSpriteScreenEntered(this js.Value, args []js.Value) interface{} {
	if callbacks.OnSpriteScreenEntered == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnSpriteScreenEntered(arg0)
	return nil
}
func gdspxOnMousePressed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnMousePressed == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	callbacks.OnMousePressed(arg0)
	return nil
}
func gdspxOnMouseReleased(this js.Value, args []js.Value) interface{} {
	if callbacks.OnMouseReleased == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	callbacks.OnMouseReleased(arg0)
	return nil
}
func gdspxOnKeyPressed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnKeyPressed == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	callbacks.OnKeyPressed(arg0)
	return nil
}
func gdspxOnKeyReleased(this js.Value, args []js.Value) interface{} {
	if callbacks.OnKeyReleased == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	callbacks.OnKeyReleased(arg0)
	return nil
}
func gdspxOnActionPressed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnActionPressed == nil {
		return nil
	}
	arg0 := JsToGdString(args[0])
	callbacks.OnActionPressed(arg0)
	return nil
}
func gdspxOnActionJustPressed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnActionJustPressed == nil {
		return nil
	}
	arg0 := JsToGdString(args[0])
	callbacks.OnActionJustPressed(arg0)
	return nil
}
func gdspxOnActionJustReleased(this js.Value, args []js.Value) interface{} {
	if callbacks.OnActionJustReleased == nil {
		return nil
	}
	arg0 := JsToGdString(args[0])
	callbacks.OnActionJustReleased(arg0)
	return nil
}
func gdspxOnAxisChanged(this js.Value, args []js.Value) interface{} {
	if callbacks.OnAxisChanged == nil {
		return nil
	}
	arg0 := JsToGdString(args[0])
	arg1 := JsToGdFloat(args[1])
	callbacks.OnAxisChanged(arg0, arg1)
	return nil
}
func gdspxOnCollisionEnter(this js.Value, args []js.Value) interface{} {
	if callbacks.OnCollisionEnter == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	arg1 := JsToGdInt(args[1])
	callbacks.OnCollisionEnter(arg0, arg1)
	return nil
}
func gdspxOnCollisionStay(this js.Value, args []js.Value) interface{} {
	if callbacks.OnCollisionStay == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	arg1 := JsToGdInt(args[1])
	callbacks.OnCollisionStay(arg0, arg1)
	return nil
}
func gdspxOnCollisionExit(this js.Value, args []js.Value) interface{} {
	if callbacks.OnCollisionExit == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	arg1 := JsToGdInt(args[1])
	callbacks.OnCollisionExit(arg0, arg1)
	return nil
}
func gdspxOnTriggerEnter(this js.Value, args []js.Value) interface{} {
	if callbacks.OnTriggerEnter == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	arg1 := JsToGdInt(args[1])
	callbacks.OnTriggerEnter(arg0, arg1)
	return nil
}
func gdspxOnTriggerStay(this js.Value, args []js.Value) interface{} {
	if callbacks.OnTriggerStay == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	arg1 := JsToGdInt(args[1])
	callbacks.OnTriggerStay(arg0, arg1)
	return nil
}
func gdspxOnTriggerExit(this js.Value, args []js.Value) interface{} {
	if callbacks.OnTriggerExit == nil {
		return nil
	}
	arg0 := JsToGdInt(args[0])
	arg1 := JsToGdInt(args[1])
	callbacks.OnTriggerExit(arg0, arg1)
	return nil
}
func gdspxOnUiReady(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiReady == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiReady(arg0)
	return nil
}
func gdspxOnUiUpdated(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiUpdated == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiUpdated(arg0)
	return nil
}
func gdspxOnUiDestroyed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiDestroyed == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiDestroyed(arg0)
	return nil
}
func gdspxOnUiPressed(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiPressed == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiPressed(arg0)
	return nil
}
func gdspxOnUiReleased(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiReleased == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiReleased(arg0)
	return nil
}
func gdspxOnUiHovered(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiHovered == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiHovered(arg0)
	return nil
}
func gdspxOnUiClicked(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiClicked == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	callbacks.OnUiClicked(arg0)
	return nil
}
func gdspxOnUiToggle(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiToggle == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	arg1 := JsToGdBool(args[1])
	callbacks.OnUiToggle(arg0, arg1)
	return nil
}
func gdspxOnUiTextChanged(this js.Value, args []js.Value) interface{} {
	if callbacks.OnUiTextChanged == nil {
		return nil
	}
	arg0 := JsToGdObj(args[0])
	arg1 := JsToGdString(args[1])
	callbacks.OnUiTextChanged(arg0, arg1)
	return nil
}
func resiterFuncPtr2Js() {
	js.Global().Set("gdspx_on_engine_start", js.FuncOf(gdspxOnEngineStart))
	js.Global().Set("gdspx_on_engine_update", js.FuncOf(gdspxOnEngineUpdate))
	js.Global().Set("gdspx_on_engine_fixed_update", js.FuncOf(gdspxOnEngineFixedUpdate))
	js.Global().Set("gdspx_on_engine_destroy", js.FuncOf(gdspxOnEngineDestroy))
	js.Global().Set("gdspx_on_scene_sprite_instantiated", js.FuncOf(gdspxOnSceneSpriteInstantiated))
	js.Global().Set("gdspx_on_sprite_ready", js.FuncOf(gdspxOnSpriteReady))
	js.Global().Set("gdspx_on_sprite_updated", js.FuncOf(gdspxOnSpriteUpdated))
	js.Global().Set("gdspx_on_sprite_fixed_updated", js.FuncOf(gdspxOnSpriteFixedUpdated))
	js.Global().Set("gdspx_on_sprite_destroyed", js.FuncOf(gdspxOnSpriteDestroyed))
	js.Global().Set("gdspx_on_sprite_frames_set_changed", js.FuncOf(gdspxOnSpriteFramesSetChanged))
	js.Global().Set("gdspx_on_sprite_animation_changed", js.FuncOf(gdspxOnSpriteAnimationChanged))
	js.Global().Set("gdspx_on_sprite_frame_changed", js.FuncOf(gdspxOnSpriteFrameChanged))
	js.Global().Set("gdspx_on_sprite_animation_looped", js.FuncOf(gdspxOnSpriteAnimationLooped))
	js.Global().Set("gdspx_on_sprite_animation_finished", js.FuncOf(gdspxOnSpriteAnimationFinished))
	js.Global().Set("gdspx_on_sprite_vfx_finished", js.FuncOf(gdspxOnSpriteVfxFinished))
	js.Global().Set("gdspx_on_sprite_screen_exited", js.FuncOf(gdspxOnSpriteScreenExited))
	js.Global().Set("gdspx_on_sprite_screen_entered", js.FuncOf(gdspxOnSpriteScreenEntered))
	js.Global().Set("gdspx_on_mouse_pressed", js.FuncOf(gdspxOnMousePressed))
	js.Global().Set("gdspx_on_mouse_released", js.FuncOf(gdspxOnMouseReleased))
	js.Global().Set("gdspx_on_key_pressed", js.FuncOf(gdspxOnKeyPressed))
	js.Global().Set("gdspx_on_key_released", js.FuncOf(gdspxOnKeyReleased))
	js.Global().Set("gdspx_on_action_pressed", js.FuncOf(gdspxOnActionPressed))
	js.Global().Set("gdspx_on_action_just_pressed", js.FuncOf(gdspxOnActionJustPressed))
	js.Global().Set("gdspx_on_action_just_released", js.FuncOf(gdspxOnActionJustReleased))
	js.Global().Set("gdspx_on_axis_changed", js.FuncOf(gdspxOnAxisChanged))
	js.Global().Set("gdspx_on_collision_enter", js.FuncOf(gdspxOnCollisionEnter))
	js.Global().Set("gdspx_on_collision_stay", js.FuncOf(gdspxOnCollisionStay))
	js.Global().Set("gdspx_on_collision_exit", js.FuncOf(gdspxOnCollisionExit))
	js.Global().Set("gdspx_on_trigger_enter", js.FuncOf(gdspxOnTriggerEnter))
	js.Global().Set("gdspx_on_trigger_stay", js.FuncOf(gdspxOnTriggerStay))
	js.Global().Set("gdspx_on_trigger_exit", js.FuncOf(gdspxOnTriggerExit))
	js.Global().Set("gdspx_on_ui_ready", js.FuncOf(gdspxOnUiReady))
	js.Global().Set("gdspx_on_ui_updated", js.FuncOf(gdspxOnUiUpdated))
	js.Global().Set("gdspx_on_ui_destroyed", js.FuncOf(gdspxOnUiDestroyed))
	js.Global().Set("gdspx_on_ui_pressed", js.FuncOf(gdspxOnUiPressed))
	js.Global().Set("gdspx_on_ui_released", js.FuncOf(gdspxOnUiReleased))
	js.Global().Set("gdspx_on_ui_hovered", js.FuncOf(gdspxOnUiHovered))
	js.Global().Set("gdspx_on_ui_clicked", js.FuncOf(gdspxOnUiClicked))
	js.Global().Set("gdspx_on_ui_toggle", js.FuncOf(gdspxOnUiToggle))
	js.Global().Set("gdspx_on_ui_text_changed", js.FuncOf(gdspxOnUiTextChanged))
}