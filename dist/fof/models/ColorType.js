"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ColorTypeToJSON = exports.ColorTypeFromJSONTyped = exports.ColorTypeFromJSON = exports.ColorType = void 0;
/**
 * Color configuration type. This color configuration provides a visual category of entities.
 * @export
 */
exports.ColorType = {
    Red: 'Red',
    DarkRed: 'DarkRed',
    Green: 'Green',
    DarkGreen: 'DarkGreen',
    LightGreen: 'LightGreen',
    Orange: 'Orange',
    White: 'White',
    Yellow: 'Yellow',
    DarkYellow: 'DarkYellow',
    Purple: 'Purple',
    Brown: 'Brown',
    Gray: 'Gray',
    Aqua: 'Aqua',
    Chocolate: 'Chocolate',
    Blue: 'Blue',
    LightBlue: 'LightBlue',
    DarkBlue: 'DarkBlue',
    Cyan: 'Cyan',
    DarkCyan: 'DarkCyan',
    Magenta: 'Magenta',
    DarkMagenta: 'DarkMagenta',
    Black: 'Black'
};
function ColorTypeFromJSON(json) {
    return ColorTypeFromJSONTyped(json, false);
}
exports.ColorTypeFromJSON = ColorTypeFromJSON;
function ColorTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.ColorTypeFromJSONTyped = ColorTypeFromJSONTyped;
function ColorTypeToJSON(value) {
    return value;
}
exports.ColorTypeToJSON = ColorTypeToJSON;
