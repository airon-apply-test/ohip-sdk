/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for storing an email.
 * @export
 * @interface EmailMessageType
 */
export interface EmailMessageType {
    /**
     * Used for Character Strings, length 0 to 80.
     * @type {string}
     * @memberof EmailMessageType
     */
    blockId?: string;
    /**
     * Email address
     * @type {string}
     * @memberof EmailMessageType
     */
    fromAddress?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof EmailMessageType
     */
    toAddress?: Array<string>;
    /**
     * Used for Character Strings, length 0 to 200.
     * @type {string}
     * @memberof EmailMessageType
     */
    subject?: string;
    /**
     * Email in base64 binary format.
     * @type {string}
     * @memberof EmailMessageType
     */
    emailBody?: string;
    /**
     * Used for Character Strings, length 0 to 100.
     * @type {string}
     * @memberof EmailMessageType
     */
    messageId?: string;
    /**
     * The date the email was received.
     * @type {Date}
     * @memberof EmailMessageType
     */
    emailReceiveDate?: Date;
    /**
     * Indicates whether the email was sent with an attachment (true) or not (false).
     * @type {boolean}
     * @memberof EmailMessageType
     */
    hasAttachment?: boolean;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof EmailMessageType
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the EmailMessageType interface.
 */
export function instanceOfEmailMessageType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EmailMessageTypeFromJSON(json: any): EmailMessageType {
    return EmailMessageTypeFromJSONTyped(json, false);
}

export function EmailMessageTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailMessageType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blockId': !exists(json, 'blockId') ? undefined : json['blockId'],
        'fromAddress': !exists(json, 'fromAddress') ? undefined : json['fromAddress'],
        'toAddress': !exists(json, 'toAddress') ? undefined : json['toAddress'],
        'subject': !exists(json, 'subject') ? undefined : json['subject'],
        'emailBody': !exists(json, 'emailBody') ? undefined : json['emailBody'],
        'messageId': !exists(json, 'messageId') ? undefined : json['messageId'],
        'emailReceiveDate': !exists(json, 'emailReceiveDate') ? undefined : (new Date(json['emailReceiveDate'])),
        'hasAttachment': !exists(json, 'hasAttachment') ? undefined : json['hasAttachment'],
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function EmailMessageTypeToJSON(value?: EmailMessageType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blockId': value.blockId,
        'fromAddress': value.fromAddress,
        'toAddress': value.toAddress,
        'subject': value.subject,
        'emailBody': value.emailBody,
        'messageId': value.messageId,
        'emailReceiveDate': value.emailReceiveDate === undefined ? undefined : (value.emailReceiveDate.toISOString()),
        'hasAttachment': value.hasAttachment,
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

