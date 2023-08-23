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
exports.SubAllocationTypeExternalAttributesToJSON = exports.SubAllocationTypeExternalAttributesFromJSONTyped = exports.SubAllocationTypeExternalAttributesFromJSON = exports.instanceOfSubAllocationTypeExternalAttributes = void 0;
const runtime_1 = require("../runtime");
const EventTypeType_1 = require("./EventTypeType");
/**
 * Check if a given object implements the SubAllocationTypeExternalAttributes interface.
 */
function instanceOfSubAllocationTypeExternalAttributes(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfSubAllocationTypeExternalAttributes = instanceOfSubAllocationTypeExternalAttributes;
function SubAllocationTypeExternalAttributesFromJSON(json) {
    return SubAllocationTypeExternalAttributesFromJSONTyped(json, false);
}
exports.SubAllocationTypeExternalAttributesFromJSON = SubAllocationTypeExternalAttributesFromJSON;
function SubAllocationTypeExternalAttributesFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'housingProtected': !(0, runtime_1.exists)(json, 'housingProtected') ? undefined : json['housingProtected'],
        'eventType': !(0, runtime_1.exists)(json, 'eventType') ? undefined : (0, EventTypeType_1.EventTypeTypeFromJSON)(json['eventType']),
        'gIId': !(0, runtime_1.exists)(json, 'gIId') ? undefined : json['gIId'],
        'rollEndDate': !(0, runtime_1.exists)(json, 'rollEndDate') ? undefined : json['rollEndDate'],
    };
}
exports.SubAllocationTypeExternalAttributesFromJSONTyped = SubAllocationTypeExternalAttributesFromJSONTyped;
function SubAllocationTypeExternalAttributesToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'housingProtected': value.housingProtected,
        'eventType': (0, EventTypeType_1.EventTypeTypeToJSON)(value.eventType),
        'gIId': value.gIId,
        'rollEndDate': value.rollEndDate,
    };
}
exports.SubAllocationTypeExternalAttributesToJSON = SubAllocationTypeExternalAttributesToJSON;
