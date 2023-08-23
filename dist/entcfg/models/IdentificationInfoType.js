"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.IdentificationInfoTypeToJSON = exports.IdentificationInfoTypeFromJSONTyped = exports.IdentificationInfoTypeFromJSON = exports.instanceOfIdentificationInfoType = void 0;
const runtime_1 = require("../runtime");
const IdentificationType_1 = require("./IdentificationType");
/**
 * Check if a given object implements the IdentificationInfoType interface.
 */
function instanceOfIdentificationInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfIdentificationInfoType = instanceOfIdentificationInfoType;
function IdentificationInfoTypeFromJSON(json) {
    return IdentificationInfoTypeFromJSONTyped(json, false);
}
exports.IdentificationInfoTypeFromJSON = IdentificationInfoTypeFromJSON;
function IdentificationInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'identification': !(0, runtime_1.exists)(json, 'identification') ? undefined : (0, IdentificationType_1.IdentificationTypeFromJSON)(json['identification']),
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
    };
}
exports.IdentificationInfoTypeFromJSONTyped = IdentificationInfoTypeFromJSONTyped;
function IdentificationInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'identification': (0, IdentificationType_1.IdentificationTypeToJSON)(value.identification),
        'id': value.id,
        'type': value.type,
    };
}
exports.IdentificationInfoTypeToJSON = IdentificationInfoTypeToJSON;
