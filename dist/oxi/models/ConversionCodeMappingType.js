"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConversionCodeMappingTypeToJSON = exports.ConversionCodeMappingTypeFromJSONTyped = exports.ConversionCodeMappingTypeFromJSON = exports.instanceOfConversionCodeMappingType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ConversionCodeMappingType interface.
 */
function instanceOfConversionCodeMappingType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConversionCodeMappingType = instanceOfConversionCodeMappingType;
function ConversionCodeMappingTypeFromJSON(json) {
    return ConversionCodeMappingTypeFromJSONTyped(json, false);
}
exports.ConversionCodeMappingTypeFromJSON = ConversionCodeMappingTypeFromJSON;
function ConversionCodeMappingTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'interfaceId': !(0, runtime_1.exists)(json, 'interfaceId') ? undefined : json['interfaceId'],
        'conversionCode': !(0, runtime_1.exists)(json, 'conversionCode') ? undefined : json['conversionCode'],
        'uDFCode': !(0, runtime_1.exists)(json, 'uDFCode') ? undefined : json['uDFCode'],
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'operaValue': !(0, runtime_1.exists)(json, 'operaValue') ? undefined : json['operaValue'],
        'externalValue': !(0, runtime_1.exists)(json, 'externalValue') ? undefined : json['externalValue'],
        'operaColumn': !(0, runtime_1.exists)(json, 'operaColumn') ? undefined : json['operaColumn'],
        'externalFieldName': !(0, runtime_1.exists)(json, 'externalFieldName') ? undefined : json['externalFieldName'],
        'profileType': !(0, runtime_1.exists)(json, 'profileType') ? undefined : json['profileType'],
        'operaToExternalDefault': !(0, runtime_1.exists)(json, 'operaToExternalDefault') ? undefined : json['operaToExternalDefault'],
        'externalToOperaDefault': !(0, runtime_1.exists)(json, 'externalToOperaDefault') ? undefined : json['externalToOperaDefault'],
        'overrideExternalDefault': !(0, runtime_1.exists)(json, 'overrideExternalDefault') ? undefined : json['overrideExternalDefault'],
        'overrideOperaDefault': !(0, runtime_1.exists)(json, 'overrideOperaDefault') ? undefined : json['overrideOperaDefault'],
        'active': !(0, runtime_1.exists)(json, 'active') ? undefined : json['active'],
        'iFCCreated': !(0, runtime_1.exists)(json, 'iFCCreated') ? undefined : json['iFCCreated'],
        'masterValue': !(0, runtime_1.exists)(json, 'masterValue') ? undefined : json['masterValue'],
    };
}
exports.ConversionCodeMappingTypeFromJSONTyped = ConversionCodeMappingTypeFromJSONTyped;
function ConversionCodeMappingTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'interfaceId': value.interfaceId,
        'conversionCode': value.conversionCode,
        'uDFCode': value.uDFCode,
        'id': value.id,
        'operaValue': value.operaValue,
        'externalValue': value.externalValue,
        'operaColumn': value.operaColumn,
        'externalFieldName': value.externalFieldName,
        'profileType': value.profileType,
        'operaToExternalDefault': value.operaToExternalDefault,
        'externalToOperaDefault': value.externalToOperaDefault,
        'overrideExternalDefault': value.overrideExternalDefault,
        'overrideOperaDefault': value.overrideOperaDefault,
        'active': value.active,
        'iFCCreated': value.iFCCreated,
        'masterValue': value.masterValue,
    };
}
exports.ConversionCodeMappingTypeToJSON = ConversionCodeMappingTypeToJSON;
