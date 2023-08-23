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
exports.AutoSettleCompFoliosToJSON = exports.AutoSettleCompFoliosFromJSONTyped = exports.AutoSettleCompFoliosFromJSON = exports.instanceOfAutoSettleCompFolios = void 0;
const runtime_1 = require("../runtime");
const AutoSettleCompFoliosCriteriaType_1 = require("./AutoSettleCompFoliosCriteriaType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the AutoSettleCompFolios interface.
 */
function instanceOfAutoSettleCompFolios(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAutoSettleCompFolios = instanceOfAutoSettleCompFolios;
function AutoSettleCompFoliosFromJSON(json) {
    return AutoSettleCompFoliosFromJSONTyped(json, false);
}
exports.AutoSettleCompFoliosFromJSON = AutoSettleCompFoliosFromJSON;
function AutoSettleCompFoliosFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'criteria': !(0, runtime_1.exists)(json, 'criteria') ? undefined : (0, AutoSettleCompFoliosCriteriaType_1.AutoSettleCompFoliosCriteriaTypeFromJSON)(json['criteria']),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.AutoSettleCompFoliosFromJSONTyped = AutoSettleCompFoliosFromJSONTyped;
function AutoSettleCompFoliosToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'criteria': (0, AutoSettleCompFoliosCriteriaType_1.AutoSettleCompFoliosCriteriaTypeToJSON)(value.criteria),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.AutoSettleCompFoliosToJSON = AutoSettleCompFoliosToJSON;
