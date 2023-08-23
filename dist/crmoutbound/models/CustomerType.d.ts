/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AlienInfoType } from './AlienInfoType';
import type { AnonymizationType } from './AnonymizationType';
import type { CountryNameType } from './CountryNameType';
import type { CustomerTypeIdentifications } from './CustomerTypeIdentifications';
import type { PersonNameType } from './PersonNameType';
/**
 * Contains basic data on the customer's identity, location, relationships, finances, memberships, etc.
 * @export
 * @interface CustomerType
 */
export interface CustomerType {
    /**
     * Detailed name information for the customer.
     * @type {Array<PersonNameType>}
     * @memberof CustomerType
     */
    personName?: Array<PersonNameType>;
    /**
     *
     * @type {AnonymizationType}
     * @memberof CustomerType
     */
    anonymization?: AnonymizationType;
    /**
     *
     * @type {CountryNameType}
     * @memberof CustomerType
     */
    citizenCountry?: CountryNameType;
    /**
     *
     * @type {CustomerTypeIdentifications}
     * @memberof CustomerType
     */
    identifications?: CustomerTypeIdentifications;
    /**
     * Profession of a person.
     * @type {string}
     * @memberof CustomerType
     */
    profession?: string;
    /**
     *
     * @type {AlienInfoType}
     * @memberof CustomerType
     */
    alienInfo?: AlienInfoType;
    /**
     *
     * @type {CountryNameType}
     * @memberof CustomerType
     */
    birthCountry?: CountryNameType;
    /**
     * Business Title.
     * @type {string}
     * @memberof CustomerType
     */
    businessTitle?: string;
    /**
     * Identifies the profile gender code selected from Gender types List of values. Gender types LOV provides the values configured at gender configuration.
     * @type {string}
     * @memberof CustomerType
     */
    gender?: string;
    /**
     * Indicates the date of birth as indicated in the document, in ISO 8601 prescribed format.
     * @type {Date}
     * @memberof CustomerType
     */
    birthDate?: Date;
    /**
     * Indicates the date of birth as masked.
     * @type {string}
     * @memberof CustomerType
     */
    birthDateMasked?: string;
    /**
     * Language identification.
     * @type {string}
     * @memberof CustomerType
     */
    language?: string;
    /**
     * Nationality code identification
     * @type {string}
     * @memberof CustomerType
     */
    nationality?: string;
    /**
     * Nationality code description
     * @type {string}
     * @memberof CustomerType
     */
    nationalityDescription?: string;
    /**
     * The supplier's ranking of the customer (e.g., VIP, numerical ranking).
     * @type {string}
     * @memberof CustomerType
     */
    customerValue?: string;
    /**
     * Credit Rating of the customer.
     * @type {string}
     * @memberof CustomerType
     */
    creditRating?: string;
    /**
     * VIP status of the customer.
     * @type {string}
     * @memberof CustomerType
     */
    vipStatus?: string;
    /**
     * Description of the VIP status.
     * @type {string}
     * @memberof CustomerType
     */
    vipDescription?: string;
    /**
     * Place of birth.
     * @type {string}
     * @memberof CustomerType
     */
    birthPlace?: string;
    /**
     * This element tells profile is property exclusive or not.
     * @type {boolean}
     * @memberof CustomerType
     */
    privateProfile?: boolean;
    /**
     * This element tells if profile is blacklisted or not.
     * @type {boolean}
     * @memberof CustomerType
     */
    blacklist?: boolean;
}
/**
 * Check if a given object implements the CustomerType interface.
 */
export declare function instanceOfCustomerType(value: object): boolean;
export declare function CustomerTypeFromJSON(json: any): CustomerType;
export declare function CustomerTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CustomerType;
export declare function CustomerTypeToJSON(value?: CustomerType | null): any;
