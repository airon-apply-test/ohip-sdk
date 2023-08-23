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

import { exists, mapValues } from '../runtime';
/**
 * Information on an email for the customer.
 * @export
 * @interface EmailType
 */
export interface EmailType {
    /**
     * Defines the e-mail address.
     * @type {string}
     * @memberof EmailType
     */
    emailAddress?: string;
    /**
     * Defines the purpose of the e-mail address (e.g. personal, business, listserve).
     * @type {string}
     * @memberof EmailType
     */
    type?: string;
    /**
     * Describes the Type code
     * @type {string}
     * @memberof EmailType
     */
    typeDescription?: string;
    /**
     * Supported Email format.
     * @type {string}
     * @memberof EmailType
     */
    emailFormat?: EmailTypeEmailFormatEnum;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof EmailType
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof EmailType
     */
    orderSequence?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof EmailType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof EmailType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof EmailType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof EmailType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof EmailType
     */
    purgeDate?: Date;
}


/**
 * @export
 */
export const EmailTypeEmailFormatEnum = {
    Html: 'Html',
    Text: 'Text'
} as const;
export type EmailTypeEmailFormatEnum = typeof EmailTypeEmailFormatEnum[keyof typeof EmailTypeEmailFormatEnum];


/**
 * Check if a given object implements the EmailType interface.
 */
export function instanceOfEmailType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EmailTypeFromJSON(json: any): EmailType {
    return EmailTypeFromJSONTyped(json, false);
}

export function EmailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'emailAddress': !exists(json, 'emailAddress') ? undefined : json['emailAddress'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'typeDescription': !exists(json, 'typeDescription') ? undefined : json['typeDescription'],
        'emailFormat': !exists(json, 'emailFormat') ? undefined : json['emailFormat'],
        'primaryInd': !exists(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !exists(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}

export function EmailTypeToJSON(value?: EmailType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'emailAddress': value.emailAddress,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'emailFormat': value.emailFormat,
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0,10)),
    };
}

