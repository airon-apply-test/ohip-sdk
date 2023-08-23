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
exports.VoidedFolioToJSON = exports.VoidedFolioFromJSONTyped = exports.VoidedFolioFromJSON = exports.instanceOfVoidedFolio = void 0;
const runtime_1 = require("../runtime");
const FolioWindowType_1 = require("./FolioWindowType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the VoidedFolio interface.
 */
function instanceOfVoidedFolio(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfVoidedFolio = instanceOfVoidedFolio;
function VoidedFolioFromJSON(json) {
    return VoidedFolioFromJSONTyped(json, false);
}
exports.VoidedFolioFromJSON = VoidedFolioFromJSON;
function VoidedFolioFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'fiscalCreditBillFolio': !(0, runtime_1.exists)(json, 'fiscalCreditBillFolio') ? undefined : (0, FolioWindowType_1.FolioWindowTypeFromJSON)(json['fiscalCreditBillFolio']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.VoidedFolioFromJSONTyped = VoidedFolioFromJSONTyped;
function VoidedFolioToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'fiscalCreditBillFolio': (0, FolioWindowType_1.FolioWindowTypeToJSON)(value.fiscalCreditBillFolio),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.VoidedFolioToJSON = VoidedFolioToJSON;
