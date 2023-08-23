"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PutRoomRelatedStatusRequestToJSON = exports.PutRoomRelatedStatusRequestFromJSONTyped = exports.PutRoomRelatedStatusRequestFromJSON = exports.instanceOfPutRoomRelatedStatusRequest = void 0;
const runtime_1 = require("../runtime");
const FrontOfficeRoomStatusType_1 = require("./FrontOfficeRoomStatusType");
const HousekeepingRoomStatusType_1 = require("./HousekeepingRoomStatusType");
const InstanceLink_1 = require("./InstanceLink");
const RoomInformationType_1 = require("./RoomInformationType");
const TurndownStatusType_1 = require("./TurndownStatusType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PutRoomRelatedStatusRequest interface.
 */
function instanceOfPutRoomRelatedStatusRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPutRoomRelatedStatusRequest = instanceOfPutRoomRelatedStatusRequest;
function PutRoomRelatedStatusRequestFromJSON(json) {
    return PutRoomRelatedStatusRequestFromJSONTyped(json, false);
}
exports.PutRoomRelatedStatusRequestFromJSON = PutRoomRelatedStatusRequestFromJSON;
function PutRoomRelatedStatusRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomList': !(0, runtime_1.exists)(json, 'roomList') ? undefined : (json['roomList'].map(RoomInformationType_1.RoomInformationTypeFromJSON)),
        'housekeepingStatus': !(0, runtime_1.exists)(json, 'housekeepingStatus') ? undefined : (0, FrontOfficeRoomStatusType_1.FrontOfficeRoomStatusTypeFromJSON)(json['housekeepingStatus']),
        'housekeepingRoomStatus': !(0, runtime_1.exists)(json, 'housekeepingRoomStatus') ? undefined : (0, HousekeepingRoomStatusType_1.HousekeepingRoomStatusTypeFromJSON)(json['housekeepingRoomStatus']),
        'turndownStatus': !(0, runtime_1.exists)(json, 'turndownStatus') ? undefined : (0, TurndownStatusType_1.TurndownStatusTypeFromJSON)(json['turndownStatus']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PutRoomRelatedStatusRequestFromJSONTyped = PutRoomRelatedStatusRequestFromJSONTyped;
function PutRoomRelatedStatusRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomList': value.roomList === undefined ? undefined : (value.roomList.map(RoomInformationType_1.RoomInformationTypeToJSON)),
        'housekeepingStatus': (0, FrontOfficeRoomStatusType_1.FrontOfficeRoomStatusTypeToJSON)(value.housekeepingStatus),
        'housekeepingRoomStatus': (0, HousekeepingRoomStatusType_1.HousekeepingRoomStatusTypeToJSON)(value.housekeepingRoomStatus),
        'turndownStatus': (0, TurndownStatusType_1.TurndownStatusTypeToJSON)(value.turndownStatus),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PutRoomRelatedStatusRequestToJSON = PutRoomRelatedStatusRequestToJSON;
