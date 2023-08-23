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
exports.ECertificateGenerationTypeToJSON = exports.ECertificateGenerationTypeFromJSONTyped = exports.ECertificateGenerationTypeFromJSON = exports.instanceOfECertificateGenerationType = void 0;
const runtime_1 = require("../runtime");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ECertificateGenerationType interface.
 */
function instanceOfECertificateGenerationType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfECertificateGenerationType = instanceOfECertificateGenerationType;
function ECertificateGenerationTypeFromJSON(json) {
    return ECertificateGenerationTypeFromJSONTyped(json, false);
}
exports.ECertificateGenerationTypeFromJSON = ECertificateGenerationTypeFromJSON;
function ECertificateGenerationTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'ruleId': !(0, runtime_1.exists)(json, 'ruleId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['ruleId']),
        'generationDetail': !(0, runtime_1.exists)(json, 'generationDetail') ? undefined : json['generationDetail'],
        'referenceValue': !(0, runtime_1.exists)(json, 'referenceValue') ? undefined : json['referenceValue'],
    };
}
exports.ECertificateGenerationTypeFromJSONTyped = ECertificateGenerationTypeFromJSONTyped;
function ECertificateGenerationTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'ruleId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.ruleId),
        'generationDetail': value.generationDetail,
        'referenceValue': value.referenceValue,
    };
}
exports.ECertificateGenerationTypeToJSON = ECertificateGenerationTypeToJSON;
