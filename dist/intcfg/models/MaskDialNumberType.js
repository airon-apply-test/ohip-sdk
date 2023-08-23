"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.MaskDialNumberTypeToJSON = exports.MaskDialNumberTypeFromJSONTyped = exports.MaskDialNumberTypeFromJSON = exports.MaskDialNumberType = void 0;
/**
 * Mask the dialled digits as per the room configuration.
 * @export
 */
exports.MaskDialNumberType = {
    NoMasking: 'NoMasking',
    MaskLast2Digits: 'MaskLast2Digits',
    MaskLast4Digits: 'MaskLast4Digits',
    Mask4DigitsLeaveLast2Intact: 'Mask4DigitsLeaveLast2Intact',
    MaskAllExceptFirst2: 'MaskAllExceptFirst2',
    MaskAll: 'MaskAll',
    MaskingPerRoom: 'MaskingPerRoom'
};
function MaskDialNumberTypeFromJSON(json) {
    return MaskDialNumberTypeFromJSONTyped(json, false);
}
exports.MaskDialNumberTypeFromJSON = MaskDialNumberTypeFromJSON;
function MaskDialNumberTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.MaskDialNumberTypeFromJSONTyped = MaskDialNumberTypeFromJSONTyped;
function MaskDialNumberTypeToJSON(value) {
    return value;
}
exports.MaskDialNumberTypeToJSON = MaskDialNumberTypeToJSON;
