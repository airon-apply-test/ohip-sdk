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
exports.PackageTransactionCodeTypeToJSON = exports.PackageTransactionCodeTypeFromJSONTyped = exports.PackageTransactionCodeTypeFromJSON = exports.instanceOfPackageTransactionCodeType = void 0;
const runtime_1 = require("../runtime");
const AmountDeterminationType_1 = require("./AmountDeterminationType");
/**
 * Check if a given object implements the PackageTransactionCodeType interface.
 */
function instanceOfPackageTransactionCodeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPackageTransactionCodeType = instanceOfPackageTransactionCodeType;
function PackageTransactionCodeTypeFromJSON(json) {
    return PackageTransactionCodeTypeFromJSONTyped(json, false);
}
exports.PackageTransactionCodeTypeFromJSON = PackageTransactionCodeTypeFromJSON;
function PackageTransactionCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : (0, AmountDeterminationType_1.AmountDeterminationTypeFromJSON)(json['type']),
    };
}
exports.PackageTransactionCodeTypeFromJSONTyped = PackageTransactionCodeTypeFromJSONTyped;
function PackageTransactionCodeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'code': value.code,
        'type': (0, AmountDeterminationType_1.AmountDeterminationTypeToJSON)(value.type),
    };
}
exports.PackageTransactionCodeTypeToJSON = PackageTransactionCodeTypeToJSON;
