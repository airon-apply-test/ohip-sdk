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
exports.ExceptionDetailTypeToJSON = exports.ExceptionDetailTypeFromJSONTyped = exports.ExceptionDetailTypeFromJSON = exports.instanceOfExceptionDetailType = void 0;
const runtime_1 = require("../runtime");
const ErrorInstance_1 = require("./ErrorInstance");
const InstanceLink_1 = require("./InstanceLink");
/**
 * Check if a given object implements the ExceptionDetailType interface.
 */
function instanceOfExceptionDetailType(value) {
    let isInstance = true;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "title" in value;
    return isInstance;
}
exports.instanceOfExceptionDetailType = instanceOfExceptionDetailType;
function ExceptionDetailTypeFromJSON(json) {
    return ExceptionDetailTypeFromJSONTyped(json, false);
}
exports.ExceptionDetailTypeFromJSON = ExceptionDetailTypeFromJSON;
function ExceptionDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'type': json['type'],
        'title': json['title'],
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : json['status'],
        'detail': !(0, runtime_1.exists)(json, 'detail') ? undefined : json['detail'],
        'instance': !(0, runtime_1.exists)(json, 'instance') ? undefined : json['instance'],
        'oerrorCode': !(0, runtime_1.exists)(json, 'o:errorCode') ? undefined : json['o:errorCode'],
        'oerrorPath': !(0, runtime_1.exists)(json, 'o:errorPath') ? undefined : json['o:errorPath'],
        'oerrorDetails': !(0, runtime_1.exists)(json, 'o:errorDetails') ? undefined : (json['o:errorDetails'].map(ErrorInstance_1.ErrorInstanceFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
    };
}
exports.ExceptionDetailTypeFromJSONTyped = ExceptionDetailTypeFromJSONTyped;
function ExceptionDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'type': value.type,
        'title': value.title,
        'status': value.status,
        'detail': value.detail,
        'instance': value.instance,
        'o:errorCode': value.oerrorCode,
        'o:errorPath': value.oerrorPath,
        'o:errorDetails': value.oerrorDetails === undefined ? undefined : (value.oerrorDetails.map(ErrorInstance_1.ErrorInstanceToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
    };
}
exports.ExceptionDetailTypeToJSON = ExceptionDetailTypeToJSON;
