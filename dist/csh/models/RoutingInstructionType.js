"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoutingInstructionTypeToJSON = exports.RoutingInstructionTypeFromJSONTyped = exports.RoutingInstructionTypeFromJSON = exports.instanceOfRoutingInstructionType = void 0;
const runtime_1 = require("../runtime");
const BillingInstructionType_1 = require("./BillingInstructionType");
const RoutingInstructionTypeDuration_1 = require("./RoutingInstructionTypeDuration");
const TrxInfoType_1 = require("./TrxInfoType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the RoutingInstructionType interface.
 */
function instanceOfRoutingInstructionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoutingInstructionType = instanceOfRoutingInstructionType;
function RoutingInstructionTypeFromJSON(json) {
    return RoutingInstructionTypeFromJSONTyped(json, false);
}
exports.RoutingInstructionTypeFromJSON = RoutingInstructionTypeFromJSON;
function RoutingInstructionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'duration': !(0, runtime_1.exists)(json, 'duration') ? undefined : (0, RoutingInstructionTypeDuration_1.RoutingInstructionTypeDurationFromJSON)(json['duration']),
        'transactionCodes': !(0, runtime_1.exists)(json, 'transactionCodes') ? undefined : (json['transactionCodes'].map(TrxInfoType_1.TrxInfoTypeFromJSON)),
        'billingInstructions': !(0, runtime_1.exists)(json, 'billingInstructions') ? undefined : (json['billingInstructions'].map(BillingInstructionType_1.BillingInstructionTypeFromJSON)),
        'creditLimit': !(0, runtime_1.exists)(json, 'creditLimit') ? undefined : json['creditLimit'],
        'percentageLimit': !(0, runtime_1.exists)(json, 'percentageLimit') ? undefined : json['percentageLimit'],
        'covers': !(0, runtime_1.exists)(json, 'covers') ? undefined : json['covers'],
        'limitUsed': !(0, runtime_1.exists)(json, 'limitUsed') ? undefined : json['limitUsed'],
        'routingLinkId': !(0, runtime_1.exists)(json, 'routingLinkId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['routingLinkId']),
    };
}
exports.RoutingInstructionTypeFromJSONTyped = RoutingInstructionTypeFromJSONTyped;
function RoutingInstructionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'duration': (0, RoutingInstructionTypeDuration_1.RoutingInstructionTypeDurationToJSON)(value.duration),
        'transactionCodes': value.transactionCodes === undefined ? undefined : (value.transactionCodes.map(TrxInfoType_1.TrxInfoTypeToJSON)),
        'billingInstructions': value.billingInstructions === undefined ? undefined : (value.billingInstructions.map(BillingInstructionType_1.BillingInstructionTypeToJSON)),
        'creditLimit': value.creditLimit,
        'percentageLimit': value.percentageLimit,
        'covers': value.covers,
        'limitUsed': value.limitUsed,
        'routingLinkId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.routingLinkId),
    };
}
exports.RoutingInstructionTypeToJSON = RoutingInstructionTypeToJSON;
