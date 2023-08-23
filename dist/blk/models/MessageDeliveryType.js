"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.MessageDeliveryTypeToJSON = exports.MessageDeliveryTypeFromJSONTyped = exports.MessageDeliveryTypeFromJSON = exports.instanceOfMessageDeliveryType = void 0;
const runtime_1 = require("../runtime");
const MessageStatusType_1 = require("./MessageStatusType");
/**
 * Check if a given object implements the MessageDeliveryType interface.
 */
function instanceOfMessageDeliveryType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMessageDeliveryType = instanceOfMessageDeliveryType;
function MessageDeliveryTypeFromJSON(json) {
    return MessageDeliveryTypeFromJSONTyped(json, false);
}
exports.MessageDeliveryTypeFromJSON = MessageDeliveryTypeFromJSON;
function MessageDeliveryTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'deliveryStatus': !(0, runtime_1.exists)(json, 'deliveryStatus') ? undefined : (0, MessageStatusType_1.MessageStatusTypeFromJSON)(json['deliveryStatus']),
        'deliveryDate': !(0, runtime_1.exists)(json, 'deliveryDate') ? undefined : (new Date(json['deliveryDate'])),
        'deliveredBy': !(0, runtime_1.exists)(json, 'deliveredBy') ? undefined : json['deliveredBy'],
        'printDate': !(0, runtime_1.exists)(json, 'printDate') ? undefined : (new Date(json['printDate'])),
        'textMessageSentDate': !(0, runtime_1.exists)(json, 'textMessageSentDate') ? undefined : (new Date(json['textMessageSentDate'])),
        'textMessageSentBy': !(0, runtime_1.exists)(json, 'textMessageSentBy') ? undefined : json['textMessageSentBy'],
        'textMessageSentById': !(0, runtime_1.exists)(json, 'textMessageSentById') ? undefined : json['textMessageSentById'],
        'textMessageRecipientNo': !(0, runtime_1.exists)(json, 'textMessageRecipientNo') ? undefined : json['textMessageRecipientNo'],
    };
}
exports.MessageDeliveryTypeFromJSONTyped = MessageDeliveryTypeFromJSONTyped;
function MessageDeliveryTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'deliveryStatus': (0, MessageStatusType_1.MessageStatusTypeToJSON)(value.deliveryStatus),
        'deliveryDate': value.deliveryDate === undefined ? undefined : (value.deliveryDate.toISOString()),
        'deliveredBy': value.deliveredBy,
        'printDate': value.printDate === undefined ? undefined : (value.printDate.toISOString()),
        'textMessageSentDate': value.textMessageSentDate === undefined ? undefined : (value.textMessageSentDate.toISOString()),
        'textMessageSentBy': value.textMessageSentBy,
        'textMessageSentById': value.textMessageSentById,
        'textMessageRecipientNo': value.textMessageRecipientNo,
    };
}
exports.MessageDeliveryTypeToJSON = MessageDeliveryTypeToJSON;
