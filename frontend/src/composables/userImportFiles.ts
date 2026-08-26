import {OpenMultipleFilesDialog,OpenDirectoryDialog} from "../../wailsjs/go/main/App";
export function importFiles() {
    console.log("importFiles")
    OpenMultipleFilesDialog().catch((err)=>{
        console.log("[js:importFiles]: ",err)
    })
}
export function importFolder() {
    console.log("importFolder")
    OpenDirectoryDialog().catch((err)=>{
        console.log("[js:importFolder]: ",err)
    })
}

