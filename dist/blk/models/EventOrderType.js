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
exports.EventOrderTypeToJSON = exports.EventOrderTypeFromJSONTyped = exports.EventOrderTypeFromJSON = exports.instanceOfEventOrderType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the EventOrderType interface.
 */
function instanceOfEventOrderType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEventOrderType = instanceOfEventOrderType;
function EventOrderTypeFromJSON(json) {
    return EventOrderTypeFromJSONTyped(json, false);
}
exports.EventOrderTypeFromJSON = EventOrderTypeFromJSON;
function EventOrderTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'distributed': !(0, runtime_1.exists)(json, 'distributed') ? undefined : json['distributed'],
        'distributedDate': !(0, runtime_1.exists)(json, 'distributedDate') ? undefined : (new Date(json['distributedDate'])),
        'distributedDateTime': !(0, runtime_1.exists)(json, 'distributedDateTime') ? undefined : (new Date(json['distributedDateTime'])),
    };
}
exports.EventOrderTypeFromJSONTyped = EventOrderTypeFromJSONTyped;
function EventOrderTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'distributed': value.distributed,
        'distributedDate': value.distributedDate === undefined ? undefined : (value.distributedDate.toISOString().substr(0, 10)),
        'distributedDateTime': value.distributedDateTime === undefined ? undefined : (value.distributedDateTime.toISOString()),
    };
}
exports.EventOrderTypeToJSON = EventOrderTypeToJSON;
