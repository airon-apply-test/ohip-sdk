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
exports.RoutingInfoTypeToJSON = exports.RoutingInfoTypeFromJSONTyped = exports.RoutingInfoTypeFromJSON = exports.instanceOfRoutingInfoType = void 0;
const runtime_1 = require("../runtime");
const RoutingInfoTypeComp_1 = require("./RoutingInfoTypeComp");
const RoutingInfoTypeFolio_1 = require("./RoutingInfoTypeFolio");
const RoutingInfoTypeRequest_1 = require("./RoutingInfoTypeRequest");
const RoutingInfoTypeRoom_1 = require("./RoutingInfoTypeRoom");
/**
 * Check if a given object implements the RoutingInfoType interface.
 */
function instanceOfRoutingInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoutingInfoType = instanceOfRoutingInfoType;
function RoutingInfoTypeFromJSON(json) {
    return RoutingInfoTypeFromJSONTyped(json, false);
}
exports.RoutingInfoTypeFromJSON = RoutingInfoTypeFromJSON;
function RoutingInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'folio': !(0, runtime_1.exists)(json, 'folio') ? undefined : (0, RoutingInfoTypeFolio_1.RoutingInfoTypeFolioFromJSON)(json['folio']),
        'room': !(0, runtime_1.exists)(json, 'room') ? undefined : (0, RoutingInfoTypeRoom_1.RoutingInfoTypeRoomFromJSON)(json['room']),
        'comp': !(0, runtime_1.exists)(json, 'comp') ? undefined : (0, RoutingInfoTypeComp_1.RoutingInfoTypeCompFromJSON)(json['comp']),
        'request': !(0, runtime_1.exists)(json, 'request') ? undefined : (0, RoutingInfoTypeRequest_1.RoutingInfoTypeRequestFromJSON)(json['request']),
        'refreshFolio': !(0, runtime_1.exists)(json, 'refreshFolio') ? undefined : json['refreshFolio'],
    };
}
exports.RoutingInfoTypeFromJSONTyped = RoutingInfoTypeFromJSONTyped;
function RoutingInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'folio': (0, RoutingInfoTypeFolio_1.RoutingInfoTypeFolioToJSON)(value.folio),
        'room': (0, RoutingInfoTypeRoom_1.RoutingInfoTypeRoomToJSON)(value.room),
        'comp': (0, RoutingInfoTypeComp_1.RoutingInfoTypeCompToJSON)(value.comp),
        'request': (0, RoutingInfoTypeRequest_1.RoutingInfoTypeRequestToJSON)(value.request),
        'refreshFolio': value.refreshFolio,
    };
}
exports.RoutingInfoTypeToJSON = RoutingInfoTypeToJSON;
