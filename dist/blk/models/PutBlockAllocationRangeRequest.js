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
exports.PutBlockAllocationRangeRequestToJSON = exports.PutBlockAllocationRangeRequestFromJSONTyped = exports.PutBlockAllocationRangeRequestFromJSON = exports.instanceOfPutBlockAllocationRangeRequest = void 0;
const runtime_1 = require("../runtime");
const BlockAllocationRangeType_1 = require("./BlockAllocationRangeType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PutBlockAllocationRangeRequest interface.
 */
function instanceOfPutBlockAllocationRangeRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPutBlockAllocationRangeRequest = instanceOfPutBlockAllocationRangeRequest;
function PutBlockAllocationRangeRequestFromJSON(json) {
    return PutBlockAllocationRangeRequestFromJSONTyped(json, false);
}
exports.PutBlockAllocationRangeRequestFromJSON = PutBlockAllocationRangeRequestFromJSON;
function PutBlockAllocationRangeRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'blockAllocationRange': !(0, runtime_1.exists)(json, 'blockAllocationRange') ? undefined : (0, BlockAllocationRangeType_1.BlockAllocationRangeTypeFromJSON)(json['blockAllocationRange']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PutBlockAllocationRangeRequestFromJSONTyped = PutBlockAllocationRangeRequestFromJSONTyped;
function PutBlockAllocationRangeRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'blockAllocationRange': (0, BlockAllocationRangeType_1.BlockAllocationRangeTypeToJSON)(value.blockAllocationRange),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PutBlockAllocationRangeRequestToJSON = PutBlockAllocationRangeRequestToJSON;
