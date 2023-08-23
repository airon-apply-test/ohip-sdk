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
exports.ChangeChannelSellLimitsByDateRangeRequestToJSON = exports.ChangeChannelSellLimitsByDateRangeRequestFromJSONTyped = exports.ChangeChannelSellLimitsByDateRangeRequestFromJSON = exports.instanceOfChangeChannelSellLimitsByDateRangeRequest = void 0;
const runtime_1 = require("../runtime");
const ChannelSellLimitScheduleType_1 = require("./ChannelSellLimitScheduleType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ChangeChannelSellLimitsByDateRangeRequest interface.
 */
function instanceOfChangeChannelSellLimitsByDateRangeRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChangeChannelSellLimitsByDateRangeRequest = instanceOfChangeChannelSellLimitsByDateRangeRequest;
function ChangeChannelSellLimitsByDateRangeRequestFromJSON(json) {
    return ChangeChannelSellLimitsByDateRangeRequestFromJSONTyped(json, false);
}
exports.ChangeChannelSellLimitsByDateRangeRequestFromJSON = ChangeChannelSellLimitsByDateRangeRequestFromJSON;
function ChangeChannelSellLimitsByDateRangeRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'sellLimits': !(0, runtime_1.exists)(json, 'sellLimits') ? undefined : (json['sellLimits'].map(ChannelSellLimitScheduleType_1.ChannelSellLimitScheduleTypeFromJSON)),
        'adjustOverlappingSchedules': !(0, runtime_1.exists)(json, 'adjustOverlappingSchedules') ? undefined : json['adjustOverlappingSchedules'],
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ChangeChannelSellLimitsByDateRangeRequestFromJSONTyped = ChangeChannelSellLimitsByDateRangeRequestFromJSONTyped;
function ChangeChannelSellLimitsByDateRangeRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'sellLimits': value.sellLimits === undefined ? undefined : (value.sellLimits.map(ChannelSellLimitScheduleType_1.ChannelSellLimitScheduleTypeToJSON)),
        'adjustOverlappingSchedules': value.adjustOverlappingSchedules,
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ChangeChannelSellLimitsByDateRangeRequestToJSON = ChangeChannelSellLimitsByDateRangeRequestToJSON;
