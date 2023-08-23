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
exports.VaultHTTPTransactionMessageTypeToJSON = exports.VaultHTTPTransactionMessageTypeFromJSONTyped = exports.VaultHTTPTransactionMessageTypeFromJSON = exports.instanceOfVaultHTTPTransactionMessageType = void 0;
const runtime_1 = require("../runtime");
const DateRangeType_1 = require("./DateRangeType");
const ErrorType_1 = require("./ErrorType");
const VaultHTTPTransactionMessageTypeAuthorizationApproval_1 = require("./VaultHTTPTransactionMessageTypeAuthorizationApproval");
const VaultHTTPTransactionType_1 = require("./VaultHTTPTransactionType");
/**
 * Check if a given object implements the VaultHTTPTransactionMessageType interface.
 */
function instanceOfVaultHTTPTransactionMessageType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfVaultHTTPTransactionMessageType = instanceOfVaultHTTPTransactionMessageType;
function VaultHTTPTransactionMessageTypeFromJSON(json) {
    return VaultHTTPTransactionMessageTypeFromJSONTyped(json, false);
}
exports.VaultHTTPTransactionMessageTypeFromJSON = VaultHTTPTransactionMessageTypeFromJSON;
function VaultHTTPTransactionMessageTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'escapedRequestContent': !(0, runtime_1.exists)(json, 'escapedRequestContent') ? undefined : json['escapedRequestContent'],
        'escapedResponseContent': !(0, runtime_1.exists)(json, 'escapedResponseContent') ? undefined : json['escapedResponseContent'],
        'hTTPTransactionDuration': !(0, runtime_1.exists)(json, 'hTTPTransactionDuration') ? undefined : (0, DateRangeType_1.DateRangeTypeFromJSON)(json['hTTPTransactionDuration']),
        'hTTPError': !(0, runtime_1.exists)(json, 'hTTPError') ? undefined : (0, ErrorType_1.ErrorTypeFromJSON)(json['hTTPError']),
        'authorizationApproval': !(0, runtime_1.exists)(json, 'authorizationApproval') ? undefined : (0, VaultHTTPTransactionMessageTypeAuthorizationApproval_1.VaultHTTPTransactionMessageTypeAuthorizationApprovalFromJSON)(json['authorizationApproval']),
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : (0, VaultHTTPTransactionType_1.VaultHTTPTransactionTypeFromJSON)(json['type']),
    };
}
exports.VaultHTTPTransactionMessageTypeFromJSONTyped = VaultHTTPTransactionMessageTypeFromJSONTyped;
function VaultHTTPTransactionMessageTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'escapedRequestContent': value.escapedRequestContent,
        'escapedResponseContent': value.escapedResponseContent,
        'hTTPTransactionDuration': (0, DateRangeType_1.DateRangeTypeToJSON)(value.hTTPTransactionDuration),
        'hTTPError': (0, ErrorType_1.ErrorTypeToJSON)(value.hTTPError),
        'authorizationApproval': (0, VaultHTTPTransactionMessageTypeAuthorizationApproval_1.VaultHTTPTransactionMessageTypeAuthorizationApprovalToJSON)(value.authorizationApproval),
        'type': (0, VaultHTTPTransactionType_1.VaultHTTPTransactionTypeToJSON)(value.type),
    };
}
exports.VaultHTTPTransactionMessageTypeToJSON = VaultHTTPTransactionMessageTypeToJSON;
