"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FetchActivityBookingsTypeToJSON = exports.FetchActivityBookingsTypeFromJSONTyped = exports.FetchActivityBookingsTypeFromJSON = exports.instanceOfFetchActivityBookingsType = void 0;
const runtime_1 = require("../runtime");
const ActivityListInner_1 = require("./ActivityListInner");
const AddressType_1 = require("./AddressType");
const PersonNameType_1 = require("./PersonNameType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the FetchActivityBookingsType interface.
 */
function instanceOfFetchActivityBookingsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFetchActivityBookingsType = instanceOfFetchActivityBookingsType;
function FetchActivityBookingsTypeFromJSON(json) {
    return FetchActivityBookingsTypeFromJSONTyped(json, false);
}
exports.FetchActivityBookingsTypeFromJSON = FetchActivityBookingsTypeFromJSON;
function FetchActivityBookingsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (json['profileId'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'personName': !(0, runtime_1.exists)(json, 'personName') ? undefined : (0, PersonNameType_1.PersonNameTypeFromJSON)(json['personName']),
        'address': !(0, runtime_1.exists)(json, 'address') ? undefined : (0, AddressType_1.AddressTypeFromJSON)(json['address']),
        'activities': !(0, runtime_1.exists)(json, 'activities') ? undefined : (json['activities'].map(ActivityListInner_1.ActivityListInnerFromJSON)),
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
    };
}
exports.FetchActivityBookingsTypeFromJSONTyped = FetchActivityBookingsTypeFromJSONTyped;
function FetchActivityBookingsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileId': value.profileId === undefined ? undefined : (value.profileId.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'personName': (0, PersonNameType_1.PersonNameTypeToJSON)(value.personName),
        'address': (0, AddressType_1.AddressTypeToJSON)(value.address),
        'activities': value.activities === undefined ? undefined : (value.activities.map(ActivityListInner_1.ActivityListInnerToJSON)),
        'hotelId': value.hotelId,
    };
}
exports.FetchActivityBookingsTypeToJSON = FetchActivityBookingsTypeToJSON;
