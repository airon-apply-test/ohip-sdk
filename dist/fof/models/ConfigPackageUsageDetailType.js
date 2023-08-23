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
exports.ConfigPackageUsageDetailTypeToJSON = exports.ConfigPackageUsageDetailTypeFromJSONTyped = exports.ConfigPackageUsageDetailTypeFromJSON = exports.instanceOfConfigPackageUsageDetailType = void 0;
const runtime_1 = require("../runtime");
const ProductSourceType_1 = require("./ProductSourceType");
/**
 * Check if a given object implements the ConfigPackageUsageDetailType interface.
 */
function instanceOfConfigPackageUsageDetailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConfigPackageUsageDetailType = instanceOfConfigPackageUsageDetailType;
function ConfigPackageUsageDetailTypeFromJSON(json) {
    return ConfigPackageUsageDetailTypeFromJSONTyped(json, false);
}
exports.ConfigPackageUsageDetailTypeFromJSON = ConfigPackageUsageDetailTypeFromJSON;
function ConfigPackageUsageDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'source': !(0, runtime_1.exists)(json, 'source') ? undefined : (0, ProductSourceType_1.ProductSourceTypeFromJSON)(json['source']),
        'usedInReservations': !(0, runtime_1.exists)(json, 'usedInReservations') ? undefined : json['usedInReservations'],
        'usedInRates': !(0, runtime_1.exists)(json, 'usedInRates') ? undefined : json['usedInRates'],
        'usedInHouseReservations': !(0, runtime_1.exists)(json, 'usedInHouseReservations') ? undefined : json['usedInHouseReservations'],
    };
}
exports.ConfigPackageUsageDetailTypeFromJSONTyped = ConfigPackageUsageDetailTypeFromJSONTyped;
function ConfigPackageUsageDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'source': (0, ProductSourceType_1.ProductSourceTypeToJSON)(value.source),
        'usedInReservations': value.usedInReservations,
        'usedInRates': value.usedInRates,
        'usedInHouseReservations': value.usedInHouseReservations,
    };
}
exports.ConfigPackageUsageDetailTypeToJSON = ConfigPackageUsageDetailTypeToJSON;
