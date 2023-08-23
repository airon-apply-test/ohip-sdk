/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Information on a telephone number for the customer.
 * @export
 * @interface TelephoneType
 */
export interface TelephoneType {
    /**
     * Indicates type of technology associated with this telephone number, such as Voice, Data, Fax, Pager, Mobile, TTY, etc.
     * @type {string}
     * @memberof TelephoneType
     */
    phoneTechType?: string;
    /**
     * Describes the type of telephone number, in the context of its general use (e.g. Home, Business, Emergency Contact, Travel Arranger, Day, Evening).
     * @type {string}
     * @memberof TelephoneType
     */
    phoneUseType?: string;
    /**
     * Description of the PhoneUseType code
     * @type {string}
     * @memberof TelephoneType
     */
    phoneUseTypeDescription?: string;
    /**
     * Telephone number assigned to a single location.
     * @type {string}
     * @memberof TelephoneType
     */
    phoneNumber?: string;
    /**
     * Extension to reach a specific party at the phone number.
     * @type {string}
     * @memberof TelephoneType
     */
    extension?: string;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof TelephoneType
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof TelephoneType
     */
    orderSequence?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof TelephoneType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof TelephoneType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof TelephoneType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof TelephoneType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof TelephoneType
     */
    purgeDate?: Date;
}

/**
 * Check if a given object implements the TelephoneType interface.
 */
export function instanceOfTelephoneType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TelephoneTypeFromJSON(json: any): TelephoneType {
    return TelephoneTypeFromJSONTyped(json, false);
}

export function TelephoneTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TelephoneType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'phoneTechType': !exists(json, 'phoneTechType') ? undefined : json['phoneTechType'],
        'phoneUseType': !exists(json, 'phoneUseType') ? undefined : json['phoneUseType'],
        'phoneUseTypeDescription': !exists(json, 'phoneUseTypeDescription') ? undefined : json['phoneUseTypeDescription'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'extension': !exists(json, 'extension') ? undefined : json['extension'],
        'primaryInd': !exists(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !exists(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}

export function TelephoneTypeToJSON(value?: TelephoneType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'phoneTechType': value.phoneTechType,
        'phoneUseType': value.phoneUseType,
        'phoneUseTypeDescription': value.phoneUseTypeDescription,
        'phoneNumber': value.phoneNumber,
        'extension': value.extension,
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0,10)),
    };
}

