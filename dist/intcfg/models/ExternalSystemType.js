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
exports.ExternalSystemTypeToJSON = exports.ExternalSystemTypeFromJSONTyped = exports.ExternalSystemTypeFromJSON = exports.instanceOfExternalSystemType = void 0;
const runtime_1 = require("../runtime");
const ExternalSystemCodeType_1 = require("./ExternalSystemCodeType");
/**
 * Check if a given object implements the ExternalSystemType interface.
 */
function instanceOfExternalSystemType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfExternalSystemType = instanceOfExternalSystemType;
function ExternalSystemTypeFromJSON(json) {
    return ExternalSystemTypeFromJSONTyped(json, false);
}
exports.ExternalSystemTypeFromJSON = ExternalSystemTypeFromJSON;
function ExternalSystemTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'externalSystemCode': !(0, runtime_1.exists)(json, 'externalSystemCode') ? undefined : (0, ExternalSystemCodeType_1.ExternalSystemCodeTypeFromJSON)(json['externalSystemCode']),
    };
}
exports.ExternalSystemTypeFromJSONTyped = ExternalSystemTypeFromJSONTyped;
function ExternalSystemTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'externalSystemCode': (0, ExternalSystemCodeType_1.ExternalSystemCodeTypeToJSON)(value.externalSystemCode),
    };
}
exports.ExternalSystemTypeToJSON = ExternalSystemTypeToJSON;
