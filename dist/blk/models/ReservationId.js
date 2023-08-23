"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationIdToJSON = exports.ReservationIdFromJSONTyped = exports.ReservationIdFromJSON = exports.instanceOfReservationId = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ReservationId interface.
 */
function instanceOfReservationId(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationId = instanceOfReservationId;
function ReservationIdFromJSON(json) {
    return ReservationIdFromJSONTyped(json, false);
}
exports.ReservationIdFromJSON = ReservationIdFromJSON;
function ReservationIdFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'url': !(0, runtime_1.exists)(json, 'url') ? undefined : json['url'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'instance': !(0, runtime_1.exists)(json, 'instance') ? undefined : json['instance'],
        'idContext': !(0, runtime_1.exists)(json, 'idContext') ? undefined : json['idContext'],
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'idExtension': !(0, runtime_1.exists)(json, 'idExtension') ? undefined : json['idExtension'],
    };
}
exports.ReservationIdFromJSONTyped = ReservationIdFromJSONTyped;
function ReservationIdToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
    };
}
exports.ReservationIdToJSON = ReservationIdToJSON;
