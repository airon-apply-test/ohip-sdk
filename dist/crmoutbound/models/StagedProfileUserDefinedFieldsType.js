"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.StagedProfileUserDefinedFieldsTypeToJSON = exports.StagedProfileUserDefinedFieldsTypeFromJSONTyped = exports.StagedProfileUserDefinedFieldsTypeFromJSON = exports.instanceOfStagedProfileUserDefinedFieldsType = void 0;
const runtime_1 = require("../runtime");
const StagedProfileCharacterUDFType_1 = require("./StagedProfileCharacterUDFType");
const StagedProfileDateUDFType_1 = require("./StagedProfileDateUDFType");
const StagedProfileNumericUDFType_1 = require("./StagedProfileNumericUDFType");
/**
 * Check if a given object implements the StagedProfileUserDefinedFieldsType interface.
 */
function instanceOfStagedProfileUserDefinedFieldsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfStagedProfileUserDefinedFieldsType = instanceOfStagedProfileUserDefinedFieldsType;
function StagedProfileUserDefinedFieldsTypeFromJSON(json) {
    return StagedProfileUserDefinedFieldsTypeFromJSONTyped(json, false);
}
exports.StagedProfileUserDefinedFieldsTypeFromJSON = StagedProfileUserDefinedFieldsTypeFromJSON;
function StagedProfileUserDefinedFieldsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'characterUDFs': !(0, runtime_1.exists)(json, 'characterUDFs') ? undefined : (json['characterUDFs'].map(StagedProfileCharacterUDFType_1.StagedProfileCharacterUDFTypeFromJSON)),
        'numericUDFs': !(0, runtime_1.exists)(json, 'numericUDFs') ? undefined : (json['numericUDFs'].map(StagedProfileNumericUDFType_1.StagedProfileNumericUDFTypeFromJSON)),
        'dateUDFs': !(0, runtime_1.exists)(json, 'dateUDFs') ? undefined : (json['dateUDFs'].map(StagedProfileDateUDFType_1.StagedProfileDateUDFTypeFromJSON)),
    };
}
exports.StagedProfileUserDefinedFieldsTypeFromJSONTyped = StagedProfileUserDefinedFieldsTypeFromJSONTyped;
function StagedProfileUserDefinedFieldsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'characterUDFs': value.characterUDFs === undefined ? undefined : (value.characterUDFs.map(StagedProfileCharacterUDFType_1.StagedProfileCharacterUDFTypeToJSON)),
        'numericUDFs': value.numericUDFs === undefined ? undefined : (value.numericUDFs.map(StagedProfileNumericUDFType_1.StagedProfileNumericUDFTypeToJSON)),
        'dateUDFs': value.dateUDFs === undefined ? undefined : (value.dateUDFs.map(StagedProfileDateUDFType_1.StagedProfileDateUDFTypeToJSON)),
    };
}
exports.StagedProfileUserDefinedFieldsTypeToJSON = StagedProfileUserDefinedFieldsTypeToJSON;
