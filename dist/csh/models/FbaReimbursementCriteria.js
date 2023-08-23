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
exports.FbaReimbursementCriteriaToJSON = exports.FbaReimbursementCriteriaFromJSONTyped = exports.FbaReimbursementCriteriaFromJSON = exports.instanceOfFbaReimbursementCriteria = void 0;
const runtime_1 = require("../runtime");
const CertificateReconciliationType_1 = require("./CertificateReconciliationType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the FbaReimbursementCriteria interface.
 */
function instanceOfFbaReimbursementCriteria(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFbaReimbursementCriteria = instanceOfFbaReimbursementCriteria;
function FbaReimbursementCriteriaFromJSON(json) {
    return FbaReimbursementCriteriaFromJSONTyped(json, false);
}
exports.FbaReimbursementCriteriaFromJSON = FbaReimbursementCriteriaFromJSON;
function FbaReimbursementCriteriaFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'certificateReconciliationTypes': !(0, runtime_1.exists)(json, 'certificateReconciliationTypes') ? undefined : (json['certificateReconciliationTypes'].map(CertificateReconciliationType_1.CertificateReconciliationTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.FbaReimbursementCriteriaFromJSONTyped = FbaReimbursementCriteriaFromJSONTyped;
function FbaReimbursementCriteriaToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'certificateReconciliationTypes': value.certificateReconciliationTypes === undefined ? undefined : (value.certificateReconciliationTypes.map(CertificateReconciliationType_1.CertificateReconciliationTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.FbaReimbursementCriteriaToJSON = FbaReimbursementCriteriaToJSON;
