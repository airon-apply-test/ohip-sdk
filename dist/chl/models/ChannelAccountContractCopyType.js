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
exports.ChannelAccountContractCopyTypeToJSON = exports.ChannelAccountContractCopyTypeFromJSONTyped = exports.ChannelAccountContractCopyTypeFromJSON = exports.instanceOfChannelAccountContractCopyType = void 0;
const runtime_1 = require("../runtime");
const ChannelAccountContractCopyDetailsType_1 = require("./ChannelAccountContractCopyDetailsType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ChannelAccountContractCopyType interface.
 */
function instanceOfChannelAccountContractCopyType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelAccountContractCopyType = instanceOfChannelAccountContractCopyType;
function ChannelAccountContractCopyTypeFromJSON(json) {
    return ChannelAccountContractCopyTypeFromJSONTyped(json, false);
}
exports.ChannelAccountContractCopyTypeFromJSON = ChannelAccountContractCopyTypeFromJSON;
function ChannelAccountContractCopyTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'targetProfileId': !(0, runtime_1.exists)(json, 'targetProfileId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['targetProfileId']),
        'targetAccountCode': !(0, runtime_1.exists)(json, 'targetAccountCode') ? undefined : json['targetAccountCode'],
        'channelAccountContractCopyDetails': !(0, runtime_1.exists)(json, 'channelAccountContractCopyDetails') ? undefined : (json['channelAccountContractCopyDetails'].map(ChannelAccountContractCopyDetailsType_1.ChannelAccountContractCopyDetailsTypeFromJSON)),
    };
}
exports.ChannelAccountContractCopyTypeFromJSONTyped = ChannelAccountContractCopyTypeFromJSONTyped;
function ChannelAccountContractCopyTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'targetProfileId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.targetProfileId),
        'targetAccountCode': value.targetAccountCode,
        'channelAccountContractCopyDetails': value.channelAccountContractCopyDetails === undefined ? undefined : (value.channelAccountContractCopyDetails.map(ChannelAccountContractCopyDetailsType_1.ChannelAccountContractCopyDetailsTypeToJSON)),
    };
}
exports.ChannelAccountContractCopyTypeToJSON = ChannelAccountContractCopyTypeToJSON;
