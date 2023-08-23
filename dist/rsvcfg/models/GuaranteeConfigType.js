"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation Master Data Management API
 * APIs to cater for Reservation Configuration in OPERA Cloud. In this module you can retrieve, create, modify or delete configuration related to Reservations, Blocks and Leisure Management.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.GuaranteeConfigTypeToJSON = exports.GuaranteeConfigTypeFromJSONTyped = exports.GuaranteeConfigTypeFromJSON = exports.instanceOfGuaranteeConfigType = void 0;
const runtime_1 = require("../runtime");
const GuaranteeConfigTypeMassCancellation_1 = require("./GuaranteeConfigTypeMassCancellation");
const GuaranteeRequirementsType_1 = require("./GuaranteeRequirementsType");
const TranslationTextType80_1 = require("./TranslationTextType80");
/**
 * Check if a given object implements the GuaranteeConfigType interface.
 */
function instanceOfGuaranteeConfigType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfGuaranteeConfigType = instanceOfGuaranteeConfigType;
function GuaranteeConfigTypeFromJSON(json) {
    return GuaranteeConfigTypeFromJSONTyped(json, false);
}
exports.GuaranteeConfigTypeFromJSON = GuaranteeConfigTypeFromJSON;
function GuaranteeConfigTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'requirements': !(0, runtime_1.exists)(json, 'requirements') ? undefined : (0, GuaranteeRequirementsType_1.GuaranteeRequirementsTypeFromJSON)(json['requirements']),
        'shortDescription': !(0, runtime_1.exists)(json, 'shortDescription') ? undefined : (0, TranslationTextType80_1.TranslationTextType80FromJSON)(json['shortDescription']),
        'paymentTypes': !(0, runtime_1.exists)(json, 'paymentTypes') ? undefined : json['paymentTypes'],
        'guaranteeCode': !(0, runtime_1.exists)(json, 'guaranteeCode') ? undefined : json['guaranteeCode'],
        'onHold': !(0, runtime_1.exists)(json, 'onHold') ? undefined : json['onHold'],
        'reserveInventory': !(0, runtime_1.exists)(json, 'reserveInventory') ? undefined : json['reserveInventory'],
        'orderSequence': !(0, runtime_1.exists)(json, 'orderSequence') ? undefined : json['orderSequence'],
        'lateArrival': !(0, runtime_1.exists)(json, 'lateArrival') ? undefined : json['lateArrival'],
        'massCancellation': !(0, runtime_1.exists)(json, 'massCancellation') ? undefined : (0, GuaranteeConfigTypeMassCancellation_1.GuaranteeConfigTypeMassCancellationFromJSON)(json['massCancellation']),
        'inactive': !(0, runtime_1.exists)(json, 'inactive') ? undefined : json['inactive'],
    };
}
exports.GuaranteeConfigTypeFromJSONTyped = GuaranteeConfigTypeFromJSONTyped;
function GuaranteeConfigTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'requirements': (0, GuaranteeRequirementsType_1.GuaranteeRequirementsTypeToJSON)(value.requirements),
        'shortDescription': (0, TranslationTextType80_1.TranslationTextType80ToJSON)(value.shortDescription),
        'paymentTypes': value.paymentTypes,
        'guaranteeCode': value.guaranteeCode,
        'onHold': value.onHold,
        'reserveInventory': value.reserveInventory,
        'orderSequence': value.orderSequence,
        'lateArrival': value.lateArrival,
        'massCancellation': (0, GuaranteeConfigTypeMassCancellation_1.GuaranteeConfigTypeMassCancellationToJSON)(value.massCancellation),
        'inactive': value.inactive,
    };
}
exports.GuaranteeConfigTypeToJSON = GuaranteeConfigTypeToJSON;
