"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.EventPrimaryInfoTypeToJSON = exports.EventPrimaryInfoTypeFromJSONTyped = exports.EventPrimaryInfoTypeFromJSON = exports.instanceOfEventPrimaryInfoType = void 0;
const runtime_1 = require("../runtime");
const EventId_1 = require("./EventId");
/**
 * Check if a given object implements the EventPrimaryInfoType interface.
 */
function instanceOfEventPrimaryInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEventPrimaryInfoType = instanceOfEventPrimaryInfoType;
function EventPrimaryInfoTypeFromJSON(json) {
    return EventPrimaryInfoTypeFromJSONTyped(json, false);
}
exports.EventPrimaryInfoTypeFromJSON = EventPrimaryInfoTypeFromJSON;
function EventPrimaryInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'eventId': !(0, runtime_1.exists)(json, 'eventId') ? undefined : (0, EventId_1.EventIdFromJSON)(json['eventId']),
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'includeSubEvents': !(0, runtime_1.exists)(json, 'includeSubEvents') ? undefined : json['includeSubEvents'],
    };
}
exports.EventPrimaryInfoTypeFromJSONTyped = EventPrimaryInfoTypeFromJSONTyped;
function EventPrimaryInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'eventId': (0, EventId_1.EventIdToJSON)(value.eventId),
        'hotelId': value.hotelId,
        'includeSubEvents': value.includeSubEvents,
    };
}
exports.EventPrimaryInfoTypeToJSON = EventPrimaryInfoTypeToJSON;
