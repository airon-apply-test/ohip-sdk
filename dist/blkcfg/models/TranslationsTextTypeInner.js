"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.TranslationsTextTypeInnerToJSON = exports.TranslationsTextTypeInnerFromJSONTyped = exports.TranslationsTextTypeInnerFromJSON = exports.instanceOfTranslationsTextTypeInner = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the TranslationsTextTypeInner interface.
 */
function instanceOfTranslationsTextTypeInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTranslationsTextTypeInner = instanceOfTranslationsTextTypeInner;
function TranslationsTextTypeInnerFromJSON(json) {
    return TranslationsTextTypeInnerFromJSONTyped(json, false);
}
exports.TranslationsTextTypeInnerFromJSON = TranslationsTextTypeInnerFromJSON;
function TranslationsTextTypeInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'value': !(0, runtime_1.exists)(json, 'value') ? undefined : json['value'],
        'language': !(0, runtime_1.exists)(json, 'language') ? undefined : json['language'],
    };
}
exports.TranslationsTextTypeInnerFromJSONTyped = TranslationsTextTypeInnerFromJSONTyped;
function TranslationsTextTypeInnerToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'value': value.value,
        'language': value.language,
    };
}
exports.TranslationsTextTypeInnerToJSON = TranslationsTextTypeInnerToJSON;
