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
exports.FixedChargeTypeToJSON = exports.FixedChargeTypeFromJSONTyped = exports.FixedChargeTypeFromJSON = exports.instanceOfFixedChargeType = void 0;
const runtime_1 = require("../runtime");
const FixedChargeDetailType_1 = require("./FixedChargeDetailType");
const FixedChargeScheduleType_1 = require("./FixedChargeScheduleType");
/**
 * Check if a given object implements the FixedChargeType interface.
 */
function instanceOfFixedChargeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFixedChargeType = instanceOfFixedChargeType;
function FixedChargeTypeFromJSON(json) {
    return FixedChargeTypeFromJSONTyped(json, false);
}
exports.FixedChargeTypeFromJSON = FixedChargeTypeFromJSON;
function FixedChargeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'schedule': !(0, runtime_1.exists)(json, 'schedule') ? undefined : (0, FixedChargeScheduleType_1.FixedChargeScheduleTypeFromJSON)(json['schedule']),
        'charge': !(0, runtime_1.exists)(json, 'charge') ? undefined : (0, FixedChargeDetailType_1.FixedChargeDetailTypeFromJSON)(json['charge']),
        'url': !(0, runtime_1.exists)(json, 'url') ? undefined : json['url'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'instance': !(0, runtime_1.exists)(json, 'instance') ? undefined : json['instance'],
        'idContext': !(0, runtime_1.exists)(json, 'idContext') ? undefined : json['idContext'],
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'idExtension': !(0, runtime_1.exists)(json, 'idExtension') ? undefined : json['idExtension'],
    };
}
exports.FixedChargeTypeFromJSONTyped = FixedChargeTypeFromJSONTyped;
function FixedChargeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'schedule': (0, FixedChargeScheduleType_1.FixedChargeScheduleTypeToJSON)(value.schedule),
        'charge': (0, FixedChargeDetailType_1.FixedChargeDetailTypeToJSON)(value.charge),
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
    };
}
exports.FixedChargeTypeToJSON = FixedChargeTypeToJSON;
