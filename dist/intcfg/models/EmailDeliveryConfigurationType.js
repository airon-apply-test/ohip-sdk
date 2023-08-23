"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.EmailDeliveryConfigurationTypeToJSON = exports.EmailDeliveryConfigurationTypeFromJSONTyped = exports.EmailDeliveryConfigurationTypeFromJSON = exports.instanceOfEmailDeliveryConfigurationType = void 0;
const runtime_1 = require("../runtime");
const EmailDeliveryFormatType_1 = require("./EmailDeliveryFormatType");
const EmailDeliveryMethodType_1 = require("./EmailDeliveryMethodType");
/**
 * Check if a given object implements the EmailDeliveryConfigurationType interface.
 */
function instanceOfEmailDeliveryConfigurationType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEmailDeliveryConfigurationType = instanceOfEmailDeliveryConfigurationType;
function EmailDeliveryConfigurationTypeFromJSON(json) {
    return EmailDeliveryConfigurationTypeFromJSONTyped(json, false);
}
exports.EmailDeliveryConfigurationTypeFromJSON = EmailDeliveryConfigurationTypeFromJSON;
function EmailDeliveryConfigurationTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'fromAddress': !(0, runtime_1.exists)(json, 'fromAddress') ? undefined : json['fromAddress'],
        'userId': !(0, runtime_1.exists)(json, 'userId') ? undefined : json['userId'],
        'userPassword': !(0, runtime_1.exists)(json, 'userPassword') ? undefined : json['userPassword'],
        'serverName': !(0, runtime_1.exists)(json, 'serverName') ? undefined : json['serverName'],
        'subject': !(0, runtime_1.exists)(json, 'subject') ? undefined : json['subject'],
        'body': !(0, runtime_1.exists)(json, 'body') ? undefined : json['body'],
        'format': !(0, runtime_1.exists)(json, 'format') ? undefined : (0, EmailDeliveryFormatType_1.EmailDeliveryFormatTypeFromJSON)(json['format']),
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : (0, EmailDeliveryMethodType_1.EmailDeliveryMethodTypeFromJSON)(json['type']),
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'activateDelivery': !(0, runtime_1.exists)(json, 'activateDelivery') ? undefined : json['activateDelivery'],
        'hTMLFormatDelivery': !(0, runtime_1.exists)(json, 'hTMLFormatDelivery') ? undefined : json['hTMLFormatDelivery'],
        'attachICalender': !(0, runtime_1.exists)(json, 'attachICalender') ? undefined : json['attachICalender'],
    };
}
exports.EmailDeliveryConfigurationTypeFromJSONTyped = EmailDeliveryConfigurationTypeFromJSONTyped;
function EmailDeliveryConfigurationTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'fromAddress': value.fromAddress,
        'userId': value.userId,
        'userPassword': value.userPassword,
        'serverName': value.serverName,
        'subject': value.subject,
        'body': value.body,
        'format': (0, EmailDeliveryFormatType_1.EmailDeliveryFormatTypeToJSON)(value.format),
        'type': (0, EmailDeliveryMethodType_1.EmailDeliveryMethodTypeToJSON)(value.type),
        'hotelId': value.hotelId,
        'activateDelivery': value.activateDelivery,
        'hTMLFormatDelivery': value.hTMLFormatDelivery,
        'attachICalender': value.attachICalender,
    };
}
exports.EmailDeliveryConfigurationTypeToJSON = EmailDeliveryConfigurationTypeToJSON;
