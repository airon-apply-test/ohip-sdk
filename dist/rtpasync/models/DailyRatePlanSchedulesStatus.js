"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * Opera Cloud Rate Plan Asynchronous Service API
 * APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.DailyRatePlanSchedulesStatusToJSON = exports.DailyRatePlanSchedulesStatusFromJSONTyped = exports.DailyRatePlanSchedulesStatusFromJSON = exports.instanceOfDailyRatePlanSchedulesStatus = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the DailyRatePlanSchedulesStatus interface.
 */
function instanceOfDailyRatePlanSchedulesStatus(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDailyRatePlanSchedulesStatus = instanceOfDailyRatePlanSchedulesStatus;
function DailyRatePlanSchedulesStatusFromJSON(json) {
    return DailyRatePlanSchedulesStatusFromJSONTyped(json, false);
}
exports.DailyRatePlanSchedulesStatusFromJSON = DailyRatePlanSchedulesStatusFromJSON;
function DailyRatePlanSchedulesStatusFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
    };
}
exports.DailyRatePlanSchedulesStatusFromJSONTyped = DailyRatePlanSchedulesStatusFromJSONTyped;
function DailyRatePlanSchedulesStatusToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
    };
}
exports.DailyRatePlanSchedulesStatusToJSON = DailyRatePlanSchedulesStatusToJSON;
