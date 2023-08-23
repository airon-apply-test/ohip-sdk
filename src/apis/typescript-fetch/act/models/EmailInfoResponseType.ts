/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AccountId } from './AccountId';
import {
    AccountIdFromJSON,
    AccountIdFromJSONTyped,
    AccountIdToJSON,
} from './AccountId';
import type { BlockId } from './BlockId';
import {
    BlockIdFromJSON,
    BlockIdFromJSONTyped,
    BlockIdToJSON,
} from './BlockId';
import type { ContactId } from './ContactId';
import {
    ContactIdFromJSON,
    ContactIdFromJSONTyped,
    ContactIdToJSON,
} from './ContactId';
import type { EmailID } from './EmailID';
import {
    EmailIDFromJSON,
    EmailIDFromJSONTyped,
    EmailIDToJSON,
} from './EmailID';

/**
 * Defines elements related to fetched external emails stored.
 * @export
 * @interface EmailInfoResponseType
 */
export interface EmailInfoResponseType {
    /**
     * 
     * @type {Array<string>}
     * @memberof EmailInfoResponseType
     */
    hotelCodes?: Array<string>;
    /**
     * 
     * @type {EmailID}
     * @memberof EmailInfoResponseType
     */
    emailID?: EmailID;
    /**
     * 
     * @type {AccountId}
     * @memberof EmailInfoResponseType
     */
    accountId?: AccountId;
    /**
     * 
     * @type {ContactId}
     * @memberof EmailInfoResponseType
     */
    contactId?: ContactId;
    /**
     * 
     * @type {BlockId}
     * @memberof EmailInfoResponseType
     */
    blockId?: BlockId;
    /**
     * Subject of the email.
     * @type {string}
     * @memberof EmailInfoResponseType
     */
    emailSubject?: string;
    /**
     * Family name or last name of the Sender.
     * @type {string}
     * @memberof EmailInfoResponseType
     */
    senderLastName?: string;
    /**
     * Given name or first name of the Sender.
     * @type {string}
     * @memberof EmailInfoResponseType
     */
    senderFirstName?: string;
    /**
     * Email address of the Sender.
     * @type {string}
     * @memberof EmailInfoResponseType
     */
    senderEmailAddress?: string;
    /**
     * Email Attachment Id for stored external emails.
     * @type {number}
     * @memberof EmailInfoResponseType
     */
    emailAttachmentId?: number;
    /**
     * Email Recieved Date Time of the Sender.
     * @type {Date}
     * @memberof EmailInfoResponseType
     */
    emailReceivedDateTime?: Date;
    /**
     * List of Email TO Recipients separated by ";"
     * @type {string}
     * @memberof EmailInfoResponseType
     */
    emailRecipients?: string;
    /**
     * Indicates if the Email body has attachments included.
     * @type {boolean}
     * @memberof EmailInfoResponseType
     */
    hasEmailAttachments?: boolean;
}

/**
 * Check if a given object implements the EmailInfoResponseType interface.
 */
export function instanceOfEmailInfoResponseType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EmailInfoResponseTypeFromJSON(json: any): EmailInfoResponseType {
    return EmailInfoResponseTypeFromJSONTyped(json, false);
}

export function EmailInfoResponseTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailInfoResponseType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelCodes': !exists(json, 'hotelCodes') ? undefined : json['hotelCodes'],
        'emailID': !exists(json, 'emailID') ? undefined : EmailIDFromJSON(json['emailID']),
        'accountId': !exists(json, 'accountId') ? undefined : AccountIdFromJSON(json['accountId']),
        'contactId': !exists(json, 'contactId') ? undefined : ContactIdFromJSON(json['contactId']),
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'emailSubject': !exists(json, 'emailSubject') ? undefined : json['emailSubject'],
        'senderLastName': !exists(json, 'senderLastName') ? undefined : json['senderLastName'],
        'senderFirstName': !exists(json, 'senderFirstName') ? undefined : json['senderFirstName'],
        'senderEmailAddress': !exists(json, 'senderEmailAddress') ? undefined : json['senderEmailAddress'],
        'emailAttachmentId': !exists(json, 'emailAttachmentId') ? undefined : json['emailAttachmentId'],
        'emailReceivedDateTime': !exists(json, 'emailReceivedDateTime') ? undefined : (new Date(json['emailReceivedDateTime'])),
        'emailRecipients': !exists(json, 'emailRecipients') ? undefined : json['emailRecipients'],
        'hasEmailAttachments': !exists(json, 'hasEmailAttachments') ? undefined : json['hasEmailAttachments'],
    };
}

export function EmailInfoResponseTypeToJSON(value?: EmailInfoResponseType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelCodes': value.hotelCodes,
        'emailID': EmailIDToJSON(value.emailID),
        'accountId': AccountIdToJSON(value.accountId),
        'contactId': ContactIdToJSON(value.contactId),
        'blockId': BlockIdToJSON(value.blockId),
        'emailSubject': value.emailSubject,
        'senderLastName': value.senderLastName,
        'senderFirstName': value.senderFirstName,
        'senderEmailAddress': value.senderEmailAddress,
        'emailAttachmentId': value.emailAttachmentId,
        'emailReceivedDateTime': value.emailReceivedDateTime === undefined ? undefined : (value.emailReceivedDateTime.toISOString()),
        'emailRecipients': value.emailRecipients,
        'hasEmailAttachments': value.hasEmailAttachments,
    };
}

