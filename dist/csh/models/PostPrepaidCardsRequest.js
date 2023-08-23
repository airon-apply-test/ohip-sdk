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
exports.PostPrepaidCardsRequestToJSON = exports.PostPrepaidCardsRequestFromJSONTyped = exports.PostPrepaidCardsRequestFromJSON = exports.instanceOfPostPrepaidCardsRequest = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const PrepaidCardCriteriaType_1 = require("./PrepaidCardCriteriaType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostPrepaidCardsRequest interface.
 */
function instanceOfPostPrepaidCardsRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostPrepaidCardsRequest = instanceOfPostPrepaidCardsRequest;
function PostPrepaidCardsRequestFromJSON(json) {
    return PostPrepaidCardsRequestFromJSONTyped(json, false);
}
exports.PostPrepaidCardsRequestFromJSON = PostPrepaidCardsRequestFromJSON;
function PostPrepaidCardsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'criteria': !(0, runtime_1.exists)(json, 'criteria') ? undefined : (0, PrepaidCardCriteriaType_1.PrepaidCardCriteriaTypeFromJSON)(json['criteria']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostPrepaidCardsRequestFromJSONTyped = PostPrepaidCardsRequestFromJSONTyped;
function PostPrepaidCardsRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'criteria': (0, PrepaidCardCriteriaType_1.PrepaidCardCriteriaTypeToJSON)(value.criteria),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostPrepaidCardsRequestToJSON = PostPrepaidCardsRequestToJSON;
