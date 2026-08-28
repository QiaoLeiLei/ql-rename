import {reactive} from "vue";
import {main} from "../wailsjs/go/models";
import {EventsOn} from "../wailsjs/runtime";
import {GetPreview,GetRules} from "../wailsjs/go/main/App";

export const dataCenter = reactive({
  rules: {} as main.Rules,
  preview: [] as Array<main.RenamePreview>,
});

EventsOn(main.EventName.DATA_UPDATE, () => {
  console.log("DATA_UPDATE")
  refreshPreview().catch((err) => console.error(err))
})

export async function initRules(): Promise<void> {
  dataCenter.rules = await GetRules()
}

async function refreshPreview(): Promise<void> {
  dataCenter.preview = await GetPreview()
}

