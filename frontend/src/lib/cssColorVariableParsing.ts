function componentToHex(c: number) {
    var hex = c.toString(16);
    return hex.length == 1 ? "0" + hex : hex;
}

function rgbToHex(r: number, g: number, b: number) {
    return "#" + componentToHex(r) + componentToHex(g) + componentToHex(b);
}

export function cssVarToHexString(cssVarString: string) {
    var cssVar = getComputedStyle(document.body).getPropertyValue(cssVarString);
    var components = cssVar.split(' ').map(str => { return parseInt(str, 10) });
    return rgbToHex(components[0], components[1], components[2])
}