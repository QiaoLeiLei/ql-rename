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
  refreshData().catch((err) => console.error(err))
})

async function refreshData(): Promise<void> {
  dataCenter.preview = await GetPreview()
  dataCenter.rules = await GetRules()
  return Promise.resolve()
}

