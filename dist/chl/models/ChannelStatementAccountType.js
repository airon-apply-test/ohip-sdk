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
exports.ChannelStatementAccountTypeToJSON = exports.ChannelStatementAccountTypeFromJSONTyped = exports.ChannelStatementAccountTypeFromJSON = exports.instanceOfChannelStatementAccountType = void 0;
const runtime_1 = require("../runtime");
const ChannelStatementDetailType_1 = require("./ChannelStatementDetailType");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ChannelStatementAccountType interface.
 */
function instanceOfChannelStatementAccountType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelStatementAccountType = instanceOfChannelStatementAccountType;
function ChannelStatementAccountTypeFromJSON(json) {
    return ChannelStatementAccountTypeFromJSONTyped(json, false);
}
exports.ChannelStatementAccountTypeFromJSON = ChannelStatementAccountTypeFromJSON;
function ChannelStatementAccountTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['profileId']),
        'accountCode': !(0, runtime_1.exists)(json, 'accountCode') ? undefined : json['accountCode'],
        'contractId': !(0, runtime_1.exists)(json, 'contractId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['contractId']),
        'beginDate': !(0, runtime_1.exists)(json, 'beginDate') ? undefined : (new Date(json['beginDate'])),
        'endDate': !(0, runtime_1.exists)(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'note': !(0, runtime_1.exists)(json, 'note') ? undefined : json['note'],
        'totalDetailsAmount': !(0, runtime_1.exists)(json, 'totalDetailsAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['totalDetailsAmount']),
        'channelAccountStatementDetails': !(0, runtime_1.exists)(json, 'channelAccountStatementDetails') ? undefined : (json['channelAccountStatementDetails'].map(ChannelStatementDetailType_1.ChannelStatementDetailTypeFromJSON)),
    };
}
exports.ChannelStatementAccountTypeFromJSONTyped = ChannelStatementAccountTypeFromJSONTyped;
function ChannelStatementAccountTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.profileId),
        'accountCode': value.accountCode,
        'contractId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.contractId),
        'beginDate': value.beginDate === undefined ? undefined : (value.beginDate.toISOString().substr(0, 10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0, 10)),
        'note': value.note,
        'totalDetailsAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.totalDetailsAmount),
        'channelAccountStatementDetails': value.channelAccountStatementDetails === undefined ? undefined : (value.channelAccountStatementDetails.map(ChannelStatementDetailType_1.ChannelStatementDetailTypeToJSON)),
    };
}
exports.ChannelStatementAccountTypeToJSON = ChannelStatementAccountTypeToJSON;
