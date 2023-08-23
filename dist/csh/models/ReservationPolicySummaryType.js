"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationPolicySummaryTypeToJSON = exports.ReservationPolicySummaryTypeFromJSONTyped = exports.ReservationPolicySummaryTypeFromJSON = exports.instanceOfReservationPolicySummaryType = void 0;
const runtime_1 = require("../runtime");
const ResCancellationPolicyType_1 = require("./ResCancellationPolicyType");
const ResDepositPolicyType_1 = require("./ResDepositPolicyType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ReservationPolicySummaryType interface.
 */
function instanceOfReservationPolicySummaryType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationPolicySummaryType = instanceOfReservationPolicySummaryType;
function ReservationPolicySummaryTypeFromJSON(json) {
    return ReservationPolicySummaryTypeFromJSONTyped(json, false);
}
exports.ReservationPolicySummaryTypeFromJSON = ReservationPolicySummaryTypeFromJSON;
function ReservationPolicySummaryTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'cancellationPolicies': !(0, runtime_1.exists)(json, 'cancellationPolicies') ? undefined : (json['cancellationPolicies'].map(ResCancellationPolicyType_1.ResCancellationPolicyTypeFromJSON)),
        'depositPolicies': !(0, runtime_1.exists)(json, 'depositPolicies') ? undefined : (json['depositPolicies'].map(ResDepositPolicyType_1.ResDepositPolicyTypeFromJSON)),
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'name': !(0, runtime_1.exists)(json, 'name') ? undefined : json['name'],
    };
}
exports.ReservationPolicySummaryTypeFromJSONTyped = ReservationPolicySummaryTypeFromJSONTyped;
function ReservationPolicySummaryTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'cancellationPolicies': value.cancellationPolicies === undefined ? undefined : (value.cancellationPolicies.map(ResCancellationPolicyType_1.ResCancellationPolicyTypeToJSON)),
        'depositPolicies': value.depositPolicies === undefined ? undefined : (value.depositPolicies.map(ResDepositPolicyType_1.ResDepositPolicyTypeToJSON)),
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'name': value.name,
    };
}
exports.ReservationPolicySummaryTypeToJSON = ReservationPolicySummaryTypeToJSON;
