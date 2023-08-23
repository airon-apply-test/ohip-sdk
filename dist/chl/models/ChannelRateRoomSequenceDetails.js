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
exports.ChannelRateRoomSequenceDetailsToJSON = exports.ChannelRateRoomSequenceDetailsFromJSONTyped = exports.ChannelRateRoomSequenceDetailsFromJSON = exports.instanceOfChannelRateRoomSequenceDetails = void 0;
const runtime_1 = require("../runtime");
const ChannelRateRoomListType_1 = require("./ChannelRateRoomListType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ChannelRateRoomSequenceDetails interface.
 */
function instanceOfChannelRateRoomSequenceDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelRateRoomSequenceDetails = instanceOfChannelRateRoomSequenceDetails;
function ChannelRateRoomSequenceDetailsFromJSON(json) {
    return ChannelRateRoomSequenceDetailsFromJSONTyped(json, false);
}
exports.ChannelRateRoomSequenceDetailsFromJSON = ChannelRateRoomSequenceDetailsFromJSON;
function ChannelRateRoomSequenceDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'channelRateRoomList': !(0, runtime_1.exists)(json, 'channelRateRoomList') ? undefined : (0, ChannelRateRoomListType_1.ChannelRateRoomListTypeFromJSON)(json['channelRateRoomList']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ChannelRateRoomSequenceDetailsFromJSONTyped = ChannelRateRoomSequenceDetailsFromJSONTyped;
function ChannelRateRoomSequenceDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'channelRateRoomList': (0, ChannelRateRoomListType_1.ChannelRateRoomListTypeToJSON)(value.channelRateRoomList),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ChannelRateRoomSequenceDetailsToJSON = ChannelRateRoomSequenceDetailsToJSON;
