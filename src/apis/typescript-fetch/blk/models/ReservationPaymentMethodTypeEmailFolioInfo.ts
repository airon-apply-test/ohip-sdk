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

import { exists, mapValues } from '../runtime';
import type { EmailType } from './EmailType';
import {
    EmailTypeFromJSON,
    EmailTypeFromJSONTyped,
    EmailTypeToJSON,
} from './EmailType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Information on an email for the customer.
 * @export
 * @interface ReservationPaymentMethodTypeEmailFolioInfo
 */
export interface ReservationPaymentMethodTypeEmailFolioInfo {
    /**
     * 
     * @type {EmailType}
     * @memberof ReservationPaymentMethodTypeEmailFolioInfo
     */
    email?: EmailType;
    /**
     * 
     * @type {boolean}
     * @memberof ReservationPaymentMethodTypeEmailFolioInfo
     */
    emailFolio?: boolean;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof ReservationPaymentMethodTypeEmailFolioInfo
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof ReservationPaymentMethodTypeEmailFolioInfo
     */
    type?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof ReservationPaymentMethodTypeEmailFolioInfo
     */
    idContext?: string;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReservationPaymentMethodTypeEmailFolioInfo
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ReservationPaymentMethodTypeEmailFolioInfo interface.
 */
export function instanceOfReservationPaymentMethodTypeEmailFolioInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationPaymentMethodTypeEmailFolioInfoFromJSON(json: any): ReservationPaymentMethodTypeEmailFolioInfo {
    return ReservationPaymentMethodTypeEmailFolioInfoFromJSONTyped(json, false);
}

export function ReservationPaymentMethodTypeEmailFolioInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationPaymentMethodTypeEmailFolioInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'email') ? undefined : EmailTypeFromJSON(json['email']),
        'emailFolio': !exists(json, 'emailFolio') ? undefined : json['emailFolio'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'idContext': !exists(json, 'idContext') ? undefined : json['idContext'],
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ReservationPaymentMethodTypeEmailFolioInfoToJSON(value?: ReservationPaymentMethodTypeEmailFolioInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email': EmailTypeToJSON(value.email),
        'emailFolio': value.emailFolio,
        'id': value.id,
        'type': value.type,
        'idContext': value.idContext,
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

