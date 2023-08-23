"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RateRoomDetailsTypeToJSON = exports.RateRoomDetailsTypeFromJSONTyped = exports.RateRoomDetailsTypeFromJSON = exports.instanceOfRateRoomDetailsType = void 0;
const runtime_1 = require("../runtime");
const MasterRestrictionStatusesType_1 = require("./MasterRestrictionStatusesType");
const MembershipSearchType_1 = require("./MembershipSearchType");
const RoomStayType_1 = require("./RoomStayType");
/**
 * Check if a given object implements the RateRoomDetailsType interface.
 */
function instanceOfRateRoomDetailsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRateRoomDetailsType = instanceOfRateRoomDetailsType;
function RateRoomDetailsTypeFromJSON(json) {
    return RateRoomDetailsTypeFromJSONTyped(json, false);
}
exports.RateRoomDetailsTypeFromJSON = RateRoomDetailsTypeFromJSON;
function RateRoomDetailsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'memberships': !(0, runtime_1.exists)(json, 'memberships') ? undefined : (json['memberships'].map(MembershipSearchType_1.MembershipSearchTypeFromJSON)),
        'restrictionType': !(0, runtime_1.exists)(json, 'restrictionType') ? undefined : (0, MasterRestrictionStatusesType_1.MasterRestrictionStatusesTypeFromJSON)(json['restrictionType']),
        'roomStays': !(0, runtime_1.exists)(json, 'roomStays') ? undefined : (json['roomStays'].map(RoomStayType_1.RoomStayTypeFromJSON)),
    };
}
exports.RateRoomDetailsTypeFromJSONTyped = RateRoomDetailsTypeFromJSONTyped;
function RateRoomDetailsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'memberships': value.memberships === undefined ? undefined : (value.memberships.map(MembershipSearchType_1.MembershipSearchTypeToJSON)),
        'restrictionType': (0, MasterRestrictionStatusesType_1.MasterRestrictionStatusesTypeToJSON)(value.restrictionType),
        'roomStays': value.roomStays === undefined ? undefined : (value.roomStays.map(RoomStayType_1.RoomStayTypeToJSON)),
    };
}
exports.RateRoomDetailsTypeToJSON = RateRoomDetailsTypeToJSON;
