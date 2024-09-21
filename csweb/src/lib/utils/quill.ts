import { QuillDeltaToHtmlConverter } from "quill-delta-to-html";

/**
 * Converts a Quill Delta string to HTML
 * @param delta - The Quill Delta string to convert
 * @returns The HTML string
 */
export const getDeltaHTML = (delta:string):string => {
    const cfg = {}

    const converter = new QuillDeltaToHtmlConverter(JSON.parse(delta || "{}").ops, cfg);
    converter.renderCustomWith(function(customOp, contextOp){
        if (customOp.insert.type === 'mention') {
            let val = customOp.insert.value;
            return `<a id="${val.id}" class="!text-primary-500 cursor-pointer">${val.value}</a>`;
        } else {
            return 'Unmanaged custom blot!';
        }
    });

    const html = converter.convert(); 

    return html;
}