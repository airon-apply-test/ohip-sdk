"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ActivityListInnerToJSON = exports.ActivityListInnerFromJSONTyped = exports.ActivityListInnerFromJSON = exports.instanceOfActivityListInner = void 0;
const runtime_1 = require("../runtime");
const ActivityListInnerDeposit_1 = require("./ActivityListInnerDeposit");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const DateTimeSpanType_1 = require("./DateTimeSpanType");
const PersonNameType_1 = require("./PersonNameType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ActivityListInner interface.
 */
function instanceOfActivityListInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfActivityListInner = instanceOfActivityListInner;
function ActivityListInnerFromJSON(json) {
    return ActivityListInnerFromJSONTyped(json, false);
}
exports.ActivityListInnerFromJSON = ActivityListInnerFromJSON;
function ActivityListInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'activityIds': !(0, runtime_1.exists)(json, 'activityIds') ? undefined : (json['activityIds'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'location': !(0, runtime_1.exists)(json, 'location') ? undefined : json['location'],
        'name': !(0, runtime_1.exists)(json, 'name') ? undefined : json['name'],
        'numberOfPersons': !(0, runtime_1.exists)(json, 'numberOfPersons') ? undefined : json['numberOfPersons'],
        'timeSpan': !(0, runtime_1.exists)(json, 'timeSpan') ? undefined : (0, DateTimeSpanType_1.DateTimeSpanTypeFromJSON)(json['timeSpan']),
        'duration': !(0, runtime_1.exists)(json, 'duration') ? undefined : (0, DateTimeSpanType_1.DateTimeSpanTypeFromJSON)(json['duration']),
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'note': !(0, runtime_1.exists)(json, 'note') ? undefined : json['note'],
        'groupCode': !(0, runtime_1.exists)(json, 'groupCode') ? undefined : json['groupCode'],
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'deposit': !(0, runtime_1.exists)(json, 'deposit') ? undefined : (0, ActivityListInnerDeposit_1.ActivityListInnerDepositFromJSON)(json['deposit']),
        'inactiveDate': !(0, runtime_1.exists)(json, 'inactiveDate') ? undefined : (new Date(json['inactiveDate'])),
        'participants': !(0, runtime_1.exists)(json, 'participants') ? undefined : (json['participants'].map(PersonNameType_1.PersonNameTypeFromJSON)),
        'uRLLink': !(0, runtime_1.exists)(json, 'uRLLink') ? undefined : json['uRLLink'],
        'extensions': !(0, runtime_1.exists)(json, 'extensions') ? undefined : json['extensions'],
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : json['status'],
        'link': !(0, runtime_1.exists)(json, 'link') ? undefined : json['link'],
    };
}
exports.ActivityListInnerFromJSONTyped = ActivityListInnerFromJSONTyped;
function ActivityListInnerToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'activityIds': value.activityIds === undefined ? undefined : (value.activityIds.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'type': value.type,
        'location': value.location,
        'name': value.name,
        'numberOfPersons': value.numberOfPersons,
        'timeSpan': (0, DateTimeSpanType_1.DateTimeSpanTypeToJSON)(value.timeSpan),
        'duration': (0, DateTimeSpanType_1.DateTimeSpanTypeToJSON)(value.duration),
        'description': value.description,
        'note': value.note,
        'groupCode': value.groupCode,
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'deposit': (0, ActivityListInnerDeposit_1.ActivityListInnerDepositToJSON)(value.deposit),
        'inactiveDate': value.inactiveDate === undefined ? undefined : (value.inactiveDate.toISOString()),
        'participants': value.participants === undefined ? undefined : (value.participants.map(PersonNameType_1.PersonNameTypeToJSON)),
        'uRLLink': value.uRLLink,
        'extensions': value.extensions,
        'status': value.status,
        'link': value.link,
    };
}
exports.ActivityListInnerToJSON = ActivityListInnerToJSON;
