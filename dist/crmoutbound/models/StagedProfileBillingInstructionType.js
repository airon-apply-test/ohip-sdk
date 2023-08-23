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
exports.StagedProfileBillingInstructionTypeToJSON = exports.StagedProfileBillingInstructionTypeFromJSONTyped = exports.StagedProfileBillingInstructionTypeFromJSON = exports.instanceOfStagedProfileBillingInstructionType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the StagedProfileBillingInstructionType interface.
 */
function instanceOfStagedProfileBillingInstructionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfStagedProfileBillingInstructionType = instanceOfStagedProfileBillingInstructionType;
function StagedProfileBillingInstructionTypeFromJSON(json) {
    return StagedProfileBillingInstructionTypeFromJSONTyped(json, false);
}
exports.StagedProfileBillingInstructionTypeFromJSON = StagedProfileBillingInstructionTypeFromJSON;
function StagedProfileBillingInstructionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'desc': !(0, runtime_1.exists)(json, 'desc') ? undefined : json['desc'],
        'routingInstructionsId': !(0, runtime_1.exists)(json, 'routingInstructionsId') ? undefined : json['routingInstructionsId'],
        'billingCode': !(0, runtime_1.exists)(json, 'billingCode') ? undefined : json['billingCode'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'errorDescription': !(0, runtime_1.exists)(json, 'errorDescription') ? undefined : json['errorDescription'],
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
    };
}
exports.StagedProfileBillingInstructionTypeFromJSONTyped = StagedProfileBillingInstructionTypeFromJSONTyped;
function StagedProfileBillingInstructionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'desc': value.desc,
        'routingInstructionsId': value.routingInstructionsId,
        'billingCode': value.billingCode,
        'hotelId': value.hotelId,
        'errorDescription': value.errorDescription,
        'id': value.id,
        'type': value.type,
    };
}
exports.StagedProfileBillingInstructionTypeToJSON = StagedProfileBillingInstructionTypeToJSON;
