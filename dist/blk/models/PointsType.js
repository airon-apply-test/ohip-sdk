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
exports.PointsTypeToJSON = exports.PointsTypeFromJSONTyped = exports.PointsTypeFromJSON = exports.instanceOfPointsType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the PointsType interface.
 */
function instanceOfPointsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPointsType = instanceOfPointsType;
function PointsTypeFromJSON(json) {
    return PointsTypeFromJSONTyped(json, false);
}
exports.PointsTypeFromJSON = PointsTypeFromJSON;
function PointsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'awardsType': !(0, runtime_1.exists)(json, 'awardsType') ? undefined : json['awardsType'],
        'points': !(0, runtime_1.exists)(json, 'points') ? undefined : json['points'],
    };
}
exports.PointsTypeFromJSONTyped = PointsTypeFromJSONTyped;
function PointsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'awardsType': value.awardsType,
        'points': value.points,
    };
}
exports.PointsTypeToJSON = PointsTypeToJSON;
