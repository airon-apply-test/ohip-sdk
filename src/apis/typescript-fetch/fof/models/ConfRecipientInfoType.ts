/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AddressInfoType } from './AddressInfoType';
import {
    AddressInfoTypeFromJSON,
    AddressInfoTypeFromJSONTyped,
    AddressInfoTypeToJSON,
} from './AddressInfoType';
import type { EmailInfoType } from './EmailInfoType';
import {
    EmailInfoTypeFromJSON,
    EmailInfoTypeFromJSONTyped,
    EmailInfoTypeToJSON,
} from './EmailInfoType';
import type { ProfileId } from './ProfileId';
import {
    ProfileIdFromJSON,
    ProfileIdFromJSONTyped,
    ProfileIdToJSON,
} from './ProfileId';
import type { ProfileNameType } from './ProfileNameType';
import {
    ProfileNameTypeFromJSON,
    ProfileNameTypeFromJSONTyped,
    ProfileNameTypeToJSON,
} from './ProfileNameType';
import type { ProfileTypeType } from './ProfileTypeType';
import {
    ProfileTypeTypeFromJSON,
    ProfileTypeTypeFromJSONTyped,
    ProfileTypeTypeToJSON,
} from './ProfileTypeType';
import type { TelephoneInfoType } from './TelephoneInfoType';
import {
    TelephoneInfoTypeFromJSON,
    TelephoneInfoTypeFromJSONTyped,
    TelephoneInfoTypeToJSON,
} from './TelephoneInfoType';

/**
 * 
 * @export
 * @interface ConfRecipientInfoType
 */
export interface ConfRecipientInfoType {
    /**
     * 
     * @type {ProfileId}
     * @memberof ConfRecipientInfoType
     */
    profileId?: ProfileId;
    /**
     * 
     * @type {ProfileNameType}
     * @memberof ConfRecipientInfoType
     */
    formerName?: ProfileNameType;
    /**
     * 
     * @type {AddressInfoType}
     * @memberof ConfRecipientInfoType
     */
    addressInfo?: AddressInfoType;
    /**
     * 
     * @type {EmailInfoType}
     * @memberof ConfRecipientInfoType
     */
    emailInfo?: EmailInfoType;
    /**
     * 
     * @type {TelephoneInfoType}
     * @memberof ConfRecipientInfoType
     */
    faxInfo?: TelephoneInfoType;
    /**
     * 
     * @type {TelephoneInfoType}
     * @memberof ConfRecipientInfoType
     */
    telephoneInfo?: TelephoneInfoType;
    /**
     * 
     * @type {ProfileTypeType}
     * @memberof ConfRecipientInfoType
     */
    recipientType?: ProfileTypeType;
}

/**
 * Check if a given object implements the ConfRecipientInfoType interface.
 */
export function instanceOfConfRecipientInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ConfRecipientInfoTypeFromJSON(json: any): ConfRecipientInfoType {
    return ConfRecipientInfoTypeFromJSONTyped(json, false);
}

export function ConfRecipientInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfRecipientInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileId': !exists(json, 'profileId') ? undefined : ProfileIdFromJSON(json['profileId']),
        'formerName': !exists(json, 'formerName') ? undefined : ProfileNameTypeFromJSON(json['formerName']),
        'addressInfo': !exists(json, 'addressInfo') ? undefined : AddressInfoTypeFromJSON(json['addressInfo']),
        'emailInfo': !exists(json, 'emailInfo') ? undefined : EmailInfoTypeFromJSON(json['emailInfo']),
        'faxInfo': !exists(json, 'faxInfo') ? undefined : TelephoneInfoTypeFromJSON(json['faxInfo']),
        'telephoneInfo': !exists(json, 'telephoneInfo') ? undefined : TelephoneInfoTypeFromJSON(json['telephoneInfo']),
        'recipientType': !exists(json, 'recipientType') ? undefined : ProfileTypeTypeFromJSON(json['recipientType']),
    };
}

export function ConfRecipientInfoTypeToJSON(value?: ConfRecipientInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileId': ProfileIdToJSON(value.profileId),
        'formerName': ProfileNameTypeToJSON(value.formerName),
        'addressInfo': AddressInfoTypeToJSON(value.addressInfo),
        'emailInfo': EmailInfoTypeToJSON(value.emailInfo),
        'faxInfo': TelephoneInfoTypeToJSON(value.faxInfo),
        'telephoneInfo': TelephoneInfoTypeToJSON(value.telephoneInfo),
        'recipientType': ProfileTypeTypeToJSON(value.recipientType),
    };
}

