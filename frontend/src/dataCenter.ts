import {reactive} from "vue";
import {main} from "../wailsjs/go/models";
import {EventsOn} from "../wailsjs/runtime";
import {Events} from "./events.gen";
import {GetPreview,GetRules} from "../wailsjs/go/main/App";

export const dataCenter = reactive({
  rules: {} as main.Rules,
  preview: [] as Array<main.RenamePreview>,
});

EventsOn(Events.EventDataUpdate, () => {
  console.log("EventDataUpdate")
  refreshPreview().catch((err) => console.error(err))
})

export async function initRules(): Promise<void> {
  dataCenter.rules = await GetRules()
}

async function refreshPreview(): Promise<void> {
  dataCenter.preview = await GetPreview()
}

