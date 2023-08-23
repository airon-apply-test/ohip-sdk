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
exports.UserDefinedFieldsTypeToJSON = exports.UserDefinedFieldsTypeFromJSONTyped = exports.UserDefinedFieldsTypeFromJSON = exports.instanceOfUserDefinedFieldsType = void 0;
const runtime_1 = require("../runtime");
const CharacterUDFType_1 = require("./CharacterUDFType");
const DateUDFType_1 = require("./DateUDFType");
const NumericUDFType_1 = require("./NumericUDFType");
/**
 * Check if a given object implements the UserDefinedFieldsType interface.
 */
function instanceOfUserDefinedFieldsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfUserDefinedFieldsType = instanceOfUserDefinedFieldsType;
function UserDefinedFieldsTypeFromJSON(json) {
    return UserDefinedFieldsTypeFromJSONTyped(json, false);
}
exports.UserDefinedFieldsTypeFromJSON = UserDefinedFieldsTypeFromJSON;
function UserDefinedFieldsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'characterUDFs': !(0, runtime_1.exists)(json, 'characterUDFs') ? undefined : (json['characterUDFs'].map(CharacterUDFType_1.CharacterUDFTypeFromJSON)),
        'numericUDFs': !(0, runtime_1.exists)(json, 'numericUDFs') ? undefined : (json['numericUDFs'].map(NumericUDFType_1.NumericUDFTypeFromJSON)),
        'dateUDFs': !(0, runtime_1.exists)(json, 'dateUDFs') ? undefined : (json['dateUDFs'].map(DateUDFType_1.DateUDFTypeFromJSON)),
    };
}
exports.UserDefinedFieldsTypeFromJSONTyped = UserDefinedFieldsTypeFromJSONTyped;
function UserDefinedFieldsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'characterUDFs': value.characterUDFs === undefined ? undefined : (value.characterUDFs.map(CharacterUDFType_1.CharacterUDFTypeToJSON)),
        'numericUDFs': value.numericUDFs === undefined ? undefined : (value.numericUDFs.map(NumericUDFType_1.NumericUDFTypeToJSON)),
        'dateUDFs': value.dateUDFs === undefined ? undefined : (value.dateUDFs.map(DateUDFType_1.DateUDFTypeToJSON)),
    };
}
exports.UserDefinedFieldsTypeToJSON = UserDefinedFieldsTypeToJSON;
