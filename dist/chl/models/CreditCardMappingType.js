"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CreditCardMappingTypeToJSON = exports.CreditCardMappingTypeFromJSONTyped = exports.CreditCardMappingTypeFromJSON = exports.instanceOfCreditCardMappingType = void 0;
const runtime_1 = require("../runtime");
const SystemInfoType_1 = require("./SystemInfoType");
const TimeSpanType_1 = require("./TimeSpanType");
/**
 * Check if a given object implements the CreditCardMappingType interface.
 */
function instanceOfCreditCardMappingType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCreditCardMappingType = instanceOfCreditCardMappingType;
function CreditCardMappingTypeFromJSON(json) {
    return CreditCardMappingTypeFromJSONTyped(json, false);
}
exports.CreditCardMappingTypeFromJSON = CreditCardMappingTypeFromJSON;
function CreditCardMappingTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'systemInfo': !(0, runtime_1.exists)(json, 'systemInfo') ? undefined : (0, SystemInfoType_1.SystemInfoTypeFromJSON)(json['systemInfo']),
        'localSystemCode': !(0, runtime_1.exists)(json, 'localSystemCode') ? undefined : json['localSystemCode'],
        'externalSystemCode': !(0, runtime_1.exists)(json, 'externalSystemCode') ? undefined : json['externalSystemCode'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'dateRange': !(0, runtime_1.exists)(json, 'dateRange') ? undefined : (0, TimeSpanType_1.TimeSpanTypeFromJSON)(json['dateRange']),
    };
}
exports.CreditCardMappingTypeFromJSONTyped = CreditCardMappingTypeFromJSONTyped;
function CreditCardMappingTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'systemInfo': (0, SystemInfoType_1.SystemInfoTypeToJSON)(value.systemInfo),
        'localSystemCode': value.localSystemCode,
        'externalSystemCode': value.externalSystemCode,
        'description': value.description,
        'dateRange': (0, TimeSpanType_1.TimeSpanTypeToJSON)(value.dateRange),
    };
}
exports.CreditCardMappingTypeToJSON = CreditCardMappingTypeToJSON;
